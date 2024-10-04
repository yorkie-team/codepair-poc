import { GetObjectCommand, PutObjectCommand, S3Client, S3ClientConfig } from "@aws-sdk/client-s3";
import { getSignedUrl } from "@aws-sdk/s3-request-presigner";
import {
	BadRequestException,
	Injectable,
	NotFoundException,
	UnauthorizedException,
	UnprocessableEntityException,
} from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { Workspace } from "@prisma/client";
import * as htmlPdf from "html-pdf-node";
import * as MarkdownIt from "markdown-it";
import { PrismaService } from "src/db/prisma.service";
import { generateRandomKey } from "src/utils/functions/random-string";
import { StorageType } from "src/utils/types/storage.type";
import { CreateUploadPresignedUrlResponse } from "./types/create-upload-url-response.type";
import { ExportFileRequestBody, ExportFileResponse } from "./types/export-file.type";

@Injectable()
export class FilesService {
	private s3Client: S3Client;
	private readonly markdown: MarkdownIt;

	constructor(
		private configService: ConfigService,
		private prismaService: PrismaService
	) {
		this.s3Client = new S3Client(this.getStorageConfig());
		this.markdown = new MarkdownIt({
			html: true,
			breaks: true,
			linkify: true,
		});
	}

	async createUploadPresignedUrl(
		workspaceId: string,
		contentLength: number,
		contentType: string
	): Promise<CreateUploadPresignedUrlResponse> {
		let workspace: Workspace;
		try {
			workspace = await this.prismaService.workspace.findFirstOrThrow({
				where: {
					id: workspaceId,
				},
			});
		} catch {
			throw new UnauthorizedException("Client unauthorized.");
		}

		if (contentLength > 10_000_000) {
			throw new UnprocessableEntityException("Content length too long.");
		}

		const fileKey = `${workspace.slug}-${generateRandomKey()}.${contentType.split("/")[1]}`;
		const command = new PutObjectCommand({
			Bucket: this.configService.get("BUCKET_NAME"),
			Key: fileKey,
			StorageClass: "INTELLIGENT_TIERING",
			ContentType: contentType,
			ContentLength: contentLength,
		});
		return {
			fileKey,
			url: await getSignedUrl(this.s3Client, command, { expiresIn: 300 }),
		};
	}

	async createDownloadPresignedUrl(fileKey: string) {
		try {
			const command = new GetObjectCommand({
				Bucket: this.configService.get("BUCKET_NAME"),
				Key: fileKey,
			});
			return getSignedUrl(this.s3Client, command, { expiresIn: 3600 });
		} catch {
			throw new NotFoundException("File not found.");
		}
	}

	async exportMarkdown(
		exportFileRequestBody: ExportFileRequestBody
	): Promise<ExportFileResponse> {
		const { exportType, content, fileName } = exportFileRequestBody;

		switch (exportType) {
			case "markdown":
				return this.exportToMarkdown(content, fileName);
			case "html":
				return this.exportToHtml(content, fileName);
			case "pdf":
				return this.exportToPdf(content, fileName);
			default:
				throw new BadRequestException("Invalid export type");
		}
	}

	private async exportToMarkdown(content: string, fileName: string): Promise<ExportFileResponse> {
		return {
			fileContent: Buffer.from(content),
			mimeType: "text/markdown",
			fileName: `${fileName}.md`,
		};
	}

	private async exportToHtml(content: string, fileName: string): Promise<ExportFileResponse> {
		const html = this.markdown.render(content);
		return {
			fileContent: Buffer.from(html),
			mimeType: "text/html",
			fileName: `${fileName}.html`,
		};
	}

	private async exportToPdf(content: string, fileName: string): Promise<ExportFileResponse> {
		const html = this.markdown.render(content);
		const options = { format: "A4" };
		const file = { content: html };

		const pdfBuffer = await htmlPdf.generatePdf(file, options);
		return {
			fileContent: pdfBuffer,
			mimeType: "application/pdf",
			fileName: `${fileName}.pdf`,
		};
	}

	private getStorageConfig = (): S3ClientConfig => {
		const bucketType: StorageType = this.configService.get("BUCKET_TYPE") || "S3";
		const region = this.configService.get("AWS_REGION") || "us-east-1";
		if (bucketType === "MINIO") {
			const endpoint = this.configService.get("MINIO_ENDPOINT");
			const accessKeyId = this.configService.get("MINIO_ACCESS_KEY");
			const secretAccessKey = this.configService.get("MINIO_SECRET_KEY");
			return {
				region,
				endpoint,
				forcePathStyle: true,
				credentials: {
					accessKeyId,
					secretAccessKey,
				},
			};
		}

		return {
			region,
		};
	};
}
