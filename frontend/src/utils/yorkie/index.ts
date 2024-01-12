import {
	ySync,
	ySyncFacet,
	YSyncConfig,
	YorkieCodeMirrorPresenceType,
	YorkieCodeMirrorDocType,
} from "./y-sync.js";
import { yRemoteSelections, yRemoteSelectionsTheme } from "./y-remote-selections.js";
import * as yorkie from "yorkie-js-sdk";

export { ySync, ySyncFacet, YSyncConfig };

export function yorkieCodeMirror<
	T extends YorkieCodeMirrorDocType,
	P extends YorkieCodeMirrorPresenceType,
>(doc: yorkie.Document<T, P>, client: yorkie.Client) {
	const ySyncConfig = new YSyncConfig(doc, client);
	const plugins = [ySyncFacet.of(ySyncConfig), ySync];

	if (client) {
		plugins.push(yRemoteSelectionsTheme, yRemoteSelections);
	}

	return plugins;
}
