"use client";

import { useRouter } from "next/navigation";
import { DocumentCard } from "@/components/documents/DocumentCard";
import { UploadDropzone } from "@/components/documents/UploadDropzone";
import { useDocuments } from "@/hooks/useDocuments";

export default function DocumentsPage() {
  const router = useRouter();

  const {
    documents,
    loading,
    loadError,
    uploading,
    uploadingName,
    searchQuery,
    setSearchQuery,
    renamingId,
    renameDraft,
    setRenameDraft,
    confirmDeleteId,
    upload,
    startRename,
    cancelRename,
    commitRename,
    startDelete,
    cancelDelete,
    confirmDelete,
  } = useDocuments();

  const showEmptyState = !loading && !uploading && documents.length === 0;

  return (
    <div className="mx-auto w-full max-w-[1180px] px-8 pb-12 pt-6">
      <div className="mb-5 flex flex-wrap items-end justify-between gap-6">
        <div>
          <div className="text-[22px] font-bold tracking-tight text-ink">
            Documents
          </div>
          <div className="mt-1 text-[13.5px] text-ink-muted">
            Upload files and let DocsAI index them for search and Q&amp;A.
          </div>
        </div>
        <input
          placeholder="Search documents..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="w-60 rounded-lg border border-border bg-surface px-3 py-2 text-[13px] text-ink outline-none"
        />
      </div>

      <UploadDropzone
        uploading={uploading}
        uploadingName={uploadingName}
        onUpload={upload}
      />

      {loadError && (
        <div className="mb-4 rounded-lg border border-danger-soft bg-danger-soft px-4 py-3 text-[13px] text-danger">
          {loadError}
        </div>
      )}

      {loading && (
        <div className="grid grid-cols-[repeat(auto-fill,minmax(230px,1fr))] gap-3.5">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <div
              key={i}
              className="rounded-[11px] border border-border bg-surface p-3.5"
            >
              <div className="h-[34px] w-[34px] animate-[docsai-skel_1.3s_linear_infinite] rounded-[9px] bg-hover" />
              <div className="mt-3.5 h-3 w-[70%] animate-[docsai-skel_1.3s_linear_infinite] rounded bg-hover" />
              <div className="mt-2 h-2.5 w-[45%] animate-[docsai-skel_1.3s_linear_infinite] rounded bg-hover" />
            </div>
          ))}
        </div>
      )}

      {showEmptyState && (
        <div className="rounded-xl border border-border bg-surface px-6 py-14 text-center">
          <div className="mx-auto flex h-10 w-10 items-center justify-center rounded-[10px] bg-hover text-lg text-ink-faint">
            ＋
          </div>
          <div className="mt-3.5 text-[14.5px] font-semibold text-ink">
            No documents yet
          </div>
          <div className="mt-1 text-[13px] text-ink-muted">
            Upload your first file above to start asking questions about it.
          </div>
        </div>
      )}

      {!loading && (
        <div className="grid grid-cols-[repeat(auto-fill,minmax(230px,1fr))] gap-3.5">
          {documents.map((doc) => (
            <DocumentCard
              key={doc.id}
              doc={doc}
              isRenaming={renamingId === doc.id}
              renameDraft={renameDraft}
              isConfirmingDelete={confirmDeleteId === doc.id}
              onOpen={() => router.push(`/documents/${doc.id}`)}
              onStartRename={() => startRename(doc)}
              onRenameDraftChange={setRenameDraft}
              onCommitRename={() => commitRename(doc.id)}
              onCancelRename={cancelRename}
              onStartDelete={() => startDelete(doc.id)}
              onCancelDelete={cancelDelete}
              onConfirmDelete={() => confirmDelete(doc.id)}
            />
          ))}
        </div>
      )}
    </div>
  );
}
