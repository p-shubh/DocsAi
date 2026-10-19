"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { StatusPill } from "@/components/StatusPill";
import { Composer } from "@/components/chat/Composer";
import { MessageBubble } from "@/components/chat/MessageBubble";
import { TypingIndicator } from "@/components/chat/TypingIndicator";
import { useDocumentChat } from "@/hooks/useDocumentChat";
import { ApiError, getDocument, getDocumentChunks, getDocumentFile } from "@/lib/api";
import { extFromMime, fmtRelTime, fmtSize } from "@/lib/format";
import type { DocumentChunk, DocumentFileResponse, DocumentResponse } from "@/lib/types";

const EXAMPLE_QUESTIONS = [
  "Summarize this document.",
  "What are the key numbers?",
  "Is there anything I should double check?",
];

export function DocumentDetailView({ documentId }: { documentId: string }) {
  const [doc, setDoc] = useState<DocumentResponse | null>(null);
  const [file, setFile] = useState<DocumentFileResponse | null>(null);
  const [chunks, setChunks] = useState<DocumentChunk[]>([]);
  const [loadError, setLoadError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    let cancelled = false;

    (async () => {
      setLoading(true);
      setLoadError(null);
      try {
        const [docRes, fileRes, chunksRes] = await Promise.all([
          getDocument(documentId),
          getDocumentFile(documentId),
          getDocumentChunks(documentId),
        ]);
        if (!cancelled) {
          setDoc(docRes);
          setFile(fileRes);
          setChunks(chunksRes.chunks);
        }
      } catch (err) {
        if (!cancelled) {
          setLoadError(err instanceof ApiError ? err.message : "Failed to load document.");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [documentId]);

  const {
    messages,
    loading: chatLoading,
    input,
    setInput,
    openSources,
    toggleSources,
    sendMessage,
    retry,
  } = useDocumentChat(documentId);

  const isEmpty = messages.length === 0 && !chatLoading;
  const sendDisabled = chatLoading || !input.trim();
  const isPdf = file?.mime_type === "application/pdf";

  const handleCopy = () => {
    const fullText = chunks.map((c) => c.text).join("\n\n");
    navigator.clipboard.writeText(fullText).then(() => {
      setCopied(true);
      setTimeout(() => setCopied(false), 1500);
    });
  };

  return (
    <div className="mx-auto flex w-full max-w-[860px] flex-1 min-h-0 flex-col px-8 pb-8 pt-6">
      <Link
        href="/documents"
        className="mb-4 inline-block w-fit text-[12.5px] font-semibold text-ink-muted"
      >
        ← Back to Documents
      </Link>

      {loading && (
        <div className="py-10 text-center text-[13px] text-ink-faint">Loading…</div>
      )}

      {!loading && loadError && (
        <div className="rounded-lg border border-danger-soft bg-danger-soft px-4 py-3 text-[13px] text-danger">
          {loadError}
        </div>
      )}

      {!loading && !loadError && doc && file && (
        <>
          <div className="flex items-center justify-between gap-3 rounded-[14px] border border-border bg-surface p-4">
            <div className="flex min-w-0 items-center gap-3">
              <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-[10px] bg-hover font-mono text-[10px] font-semibold text-ink-muted">
                {extFromMime(doc.mime_type, doc.original_name)}
              </div>
              <div className="min-w-0">
                <div className="truncate text-[15px] font-semibold text-ink">
                  {doc.original_name}
                </div>
                <div className="text-[11.5px] text-ink-faint">
                  {fmtSize(doc.size)} · {fmtRelTime(doc.uploaded_at)}
                </div>
              </div>
            </div>
            <StatusPill status={doc.status} />
          </div>

          <div className="mt-3.5">
            {isPdf ? (
              <iframe
                src={file.file_url}
                title={doc.original_name}
                className="h-[45vh] w-full rounded-[14px] border border-border"
              />
            ) : (
              <div className="rounded-[14px] border border-border bg-surface">
                <div className="flex items-center justify-between gap-3 border-b border-border px-4 py-2.5">
                  <div className="text-[12px] font-semibold text-ink-muted">
                    Extracted text{chunks.length === 0 ? " (unavailable)" : ""}
                  </div>
                  <div className="flex items-center gap-2">
                    {chunks.length > 0 && (
                      <button
                        type="button"
                        onClick={handleCopy}
                        className="rounded-md border border-border px-2.5 py-1 text-[11.5px] font-semibold text-ink-muted"
                      >
                        {copied ? "Copied!" : "Copy text"}
                      </button>
                    )}
                    <a
                      href={file.file_url}
                      target="_blank"
                      rel="noreferrer"
                      className="rounded-md bg-accent px-2.5 py-1 text-[11.5px] font-semibold text-white"
                    >
                      Download original
                    </a>
                  </div>
                </div>

                {chunks.length === 0 ? (
                  <div className="px-6 py-8 text-center text-[13px] text-ink-muted">
                    No extracted text is available yet for this document.
                  </div>
                ) : (
                  <div className="max-h-[45vh] overflow-y-auto px-4 py-3.5">
                    {chunks.map((chunk, i) => {
                      const showPageLabel =
                        i === 0 || chunks[i - 1].page !== chunk.page;
                      return (
                        <div key={chunk.chunk_index} className="mb-3.5">
                          {showPageLabel && (
                            <div className="mb-1 text-[11px] font-bold uppercase tracking-wide text-ink-faint">
                              Page {chunk.page}
                              {chunk.section ? ` · ${chunk.section}` : ""}
                            </div>
                          )}
                          <div className="whitespace-pre-wrap text-[13px] leading-relaxed text-ink select-text">
                            {chunk.text}
                          </div>
                        </div>
                      );
                    })}
                  </div>
                )}
              </div>
            )}
          </div>

          <div className="mt-6 flex flex-1 min-h-0 flex-col border-t border-border pt-5">
            <div className="text-[15px] font-bold tracking-tight text-ink">
              Chat about this document
            </div>

            <div className="mt-4 flex-1 overflow-y-auto">
              {isEmpty && (
                <div className="mx-auto mt-6 max-w-[480px] text-center">
                  <div className="flex flex-wrap justify-center gap-2">
                    {EXAMPLE_QUESTIONS.map((q) => (
                      <button
                        type="button"
                        key={q}
                        onClick={() => setInput(q)}
                        className="rounded-full border border-border bg-surface px-3 py-1.5 text-[12.5px] text-ink-muted"
                      >
                        {q}
                      </button>
                    ))}
                  </div>
                </div>
              )}

              {messages.map((msg) => (
                <MessageBubble
                  key={msg.id}
                  message={msg}
                  sourcesOpen={!!openSources[msg.id]}
                  onToggleSources={() => toggleSources(msg.id)}
                  onRetry={() => msg.kind === "error" && retry(msg.id, msg.question)}
                />
              ))}

              {chatLoading && <TypingIndicator />}
            </div>

            <Composer
              value={input}
              disabled={chatLoading}
              sendDisabled={sendDisabled}
              onChange={setInput}
              onSend={sendMessage}
            />
          </div>
        </>
      )}
    </div>
  );
}
