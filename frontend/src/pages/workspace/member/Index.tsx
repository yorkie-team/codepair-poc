import { useParams } from "react-router-dom";
import { useGetWorkspaceQuery } from "../../../hooks/api/workspace";
import {
	Box,
	Button,
	CircularProgress,
	Container,
	Divider,
	Stack,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	Typography,
} from "@mui/material";
import InfiniteScroll from "react-infinite-scroller";
import { useGetWorkspaceUserListQuery } from "../../../hooks/api/workspaceUser";
import { useMemo, useState } from "react";
import AddIcon from "@mui/icons-material/Add";
import { User } from "../../../hooks/api/types/user";
import MemberModal from "../../../components/modals/MemberModal";

function MemberIndex() {
	const params = useParams();
	const { data: workspace } = useGetWorkspaceQuery(params.workspaceSlug);
	const {
		data: workspaceUserPageList,
		fetchNextPage,
		hasNextPage,
	} = useGetWorkspaceUserListQuery(workspace?.id);
	const [memberModalOpen, setMemberModalOpen] = useState(false);

	const userList = useMemo(() => {
		return (
			workspaceUserPageList?.pages.reduce((prev, page) => {
				return prev.concat(page.workspaceUsers);
			}, [] as User[]) ?? []
		);
	}, [workspaceUserPageList?.pages]);

	const handleMemberModalOpen = () => {
		setMemberModalOpen((prev) => !prev);
	};

	return (
		<Container maxWidth="lg">
			<Stack gap={4} py={6}>
				<Stack direction="row" justifyContent="space-between" alignItems="center" px={2}>
					<Typography variant="h5" fontWeight="bold">
						{workspace?.title}{" "}
						<Typography component="span" variant="inherit" color="primary">
							{userList.length}
						</Typography>
					</Typography>
					<Button
						variant="contained"
						startIcon={<AddIcon />}
						onClick={handleMemberModalOpen}
					>
						Add Members
					</Button>
				</Stack>
				<TableContainer component={Box}>
					<Divider />
					<InfiniteScroll
						pageStart={0}
						loadMore={() => fetchNextPage()}
						hasMore={hasNextPage}
						loader={
							<Box className="loader" key={0}>
								<CircularProgress size={20} />
							</Box>
						}
						useWindow={false}
						style={{
							width: "100%",
						}}
					>
						<Table>
							<TableHead>
								<TableRow>
									<TableCell>Name</TableCell>
									<TableCell>Role</TableCell>
									<TableCell></TableCell>
								</TableRow>
							</TableHead>
							<TableBody
								style={{
									maxHeight: "100%",
									overflow: "auto",
								}}
							>
								{userList.map((row) => (
									<TableRow
										key={row.id}
										sx={{
											width: "100%",
										}}
									>
										<TableCell component="th" scope="row">
											{row.nickname}
										</TableCell>
										<TableCell>-</TableCell>
										<TableCell>-</TableCell>
									</TableRow>
								))}
							</TableBody>
						</Table>
					</InfiniteScroll>
				</TableContainer>
			</Stack>
			<MemberModal open={memberModalOpen} onClose={handleMemberModalOpen} />
		</Container>
	);
}

export default MemberIndex;
