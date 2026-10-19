"use client";

import { StatusPill } from "@/components/StatusPill";
import { extFromMime, fmtRelTime, fmtSize } from "@/lib/format";
import type { DocumentResponse } from "@/lib/types";

export function DocumentCard({
  doc,
  isRenaming,
  renameDraft,
  isConfirmingDelete,
  onOpen,
  onStartRename,
  onRenameDraftChange,
  onCommitRename,
  onCancelRename,
  onStartDelete,
  onCancelDelete,
  onConfirmDelete,
}: {
  doc: DocumentResponse;
  isRenaming: boolean;
  renameDraft: string;
  isConfirmingDelete: boolean;
  onOpen: () => void;
  onStartRename: () => void;
  onRenameDraftChange: (value: string) => void;
  onCommitRename: () => void;
  onCancelRename: () => void;
  onStartDelete: () => void;
  onCancelDelete: () => void;
  onConfirmDelete: () => void;
}) {
  return (
    <div
      onClick={onOpen}
      className="relative cursor-pointer rounded-[11px] border border-border bg-surface p-3.5 transition-colors hover:border-accent"
    >
      <div className="flex items-center justify-between">
        <div className="flex h-[34px] w-[34px] items-center justify-center rounded-[9px] bg-hover font-mono text-[9.5px] font-semibold text-ink-muted">
          {extFromMime(doc.mime_type, doc.original_name)}
        </div>
        <StatusPill status={doc.status} />
      </div>

      {isRenaming ? (
        <input
          autoFocus
          onClick={(e) => e.stopPropagation()}
          value={renameDraft}
          onChange={(e) => onRenameDraftChange(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") onCommitRename();
            if (e.key === "Escape") onCancelRename();
          }}
          onBlur={onCommitRename}
          className="mt-3 w-full rounded-md border border-accent bg-surface px-2 py-1 text-[13.5px] text-ink outline-none"
        />
      ) : (
        <div
          title={doc.original_name}
          className="mt-3 truncate text-[13.5px] font-semibold text-ink"
        >
          {doc.original_name}
        </div>
      )}

      <div className="mt-1.5 text-[11.5px] text-ink-faint">
        {fmtSize(doc.size)} · {fmtRelTime(doc.uploaded_at)}
      </div>

      <div className="mt-3 flex gap-1.5">
        <button
          type="button"
          onClick={(e) => {
            e.stopPropagation();
            onStartRename();
          }}
          className="flex-1 rounded-md border border-border px-2 py-1.5 text-[11.5px] text-ink-muted"
        >
          Rename
        </button>
        <button
          type="button"
          onClick={(e) => {
            e.stopPropagation();
            onStartDelete();
          }}
          className="flex-1 rounded-md border border-border px-2 py-1.5 text-[11.5px] text-ink-faint"
        >
          Delete
        </button>
      </div>

      {isConfirmingDelete && (
        <div
          onClick={(e) => e.stopPropagation()}
          className="absolute left-2.5 right-2.5 top-full z-10 mt-1 rounded-[10px] border border-border bg-surface p-3 shadow-lg"
        >
          <div className="mb-2 text-[12.5px] text-ink">
            Delete &ldquo;{doc.original_name}&rdquo;? This can&apos;t be undone.
          </div>
          <div className="flex justify-end gap-2">
            <button
              type="button"
              onClick={onCancelDelete}
              className="rounded-md border border-border px-2.5 py-1 text-xs text-ink-muted"
            >
              Cancel
            </button>
            <button
              type="button"
              onClick={onConfirmDelete}
              className="rounded-md bg-danger px-2.5 py-1 text-xs font-semibold text-white"
            >
              Delete
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
