import * as cmState from "@codemirror/state"; // eslint-disable-line
import * as cmView from "@codemirror/view"; // eslint-disable-line
import * as yorkie from "yorkie-js-sdk";

export type YorkieCodeMirrorDocType = {
	content: yorkie.Text<yorkie.Indexable>;
};

export type YorkieCodeMirrorPresenceType = {
	selection: yorkie.TextPosStructRange | null;
};

export class YSyncConfig<
	T extends YorkieCodeMirrorDocType,
	P extends YorkieCodeMirrorPresenceType,
> {
	doc: yorkie.Document<T, P>;
	client: yorkie.Client;
	constructor(doc: yorkie.Document<T, P>, client: yorkie.Client) {
		this.doc = doc;
		this.client = client;
	}
}

export const ySyncFacet: cmState.Facet<
	YSyncConfig<YorkieCodeMirrorDocType, YorkieCodeMirrorPresenceType>,
	YSyncConfig<YorkieCodeMirrorDocType, YorkieCodeMirrorPresenceType>
> = cmState.Facet.define({
	combine(inputs) {
		return inputs[inputs.length - 1];
	},
});

export const yorkieSyncAnnotation: cmState.AnnotationType<
	YSyncConfig<YorkieCodeMirrorDocType, YorkieCodeMirrorPresenceType>
> = cmState.Annotation.define();

class YSyncPluginValue implements cmView.PluginValue {
	view: cmView.EditorView;
	conf: YSyncConfig<YorkieCodeMirrorDocType, YorkieCodeMirrorPresenceType>;
	_doc: yorkie.Document<YorkieCodeMirrorDocType, YorkieCodeMirrorPresenceType>;
	_observer: yorkie.NextFn<yorkie.DocEvent<YorkieCodeMirrorPresenceType>>;
	_unsubscribe: yorkie.Unsubscribe;

	constructor(view: cmView.EditorView) {
		this.view = view;
		this.conf = view.state.facet(ySyncFacet);

		this._observer = (event) => {
			if (event.type !== "remote-change") return;

			const { operations } = event.value;

			operations.forEach((op) => {
				if (op.type === "edit") {
					const changes = [
						{
							from: Math.max(0, op.from),
							to: Math.max(0, op.to),
							insert: op.value!.content,
						},
					];

					view.dispatch({
						changes,
						annotations: [yorkieSyncAnnotation.of(this.conf)],
					});
				}
			});
		};
		this._doc = this.conf.doc;
		this._unsubscribe = this._doc.subscribe("$.content", this._observer);
	}

	update(update: cmView.ViewUpdate) {
		if (
			!update.docChanged ||
			(update.transactions.length > 0 &&
				update.transactions[0].annotation(yorkieSyncAnnotation) === this.conf)
		) {
			return;
		}

		this._doc.update((root, presence) => {
			update.changes.iterChanges((fromA, toA, fromB, toB, insert) => {
				if (!root.content) {
					root.content = new yorkie.Text();
				}
				const insertText = insert.sliceString(0, insert.length, "\n");
				const updatedIndexRange = root.content.edit(fromA, toA, insertText);
				if (updatedIndexRange) {
					presence.set({
						selection: root.content.indexRangeToPosRange(updatedIndexRange),
					} as unknown as YorkieCodeMirrorPresenceType);
				}
			});
		});
	}

	destroy() {
		this._unsubscribe();
	}
}

export const ySync = cmView.ViewPlugin.fromClass(YSyncPluginValue);
