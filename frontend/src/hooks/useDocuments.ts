"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import {
  ApiError,
  deleteDocument,
  listDocuments,
  renameDocument,
  uploadDocument,
} from "@/lib/api";
import type { DocumentResponse } from "@/lib/types";

const NON_TERMINAL = new Set(["uploaded", "queued", "processing"]);
const POLL_INTERVAL_MS = 3000;

export function useDocuments() {
  const [documents, setDocuments] = useState<DocumentResponse[]>([]);
  const [loading, setLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);
  const [uploading, setUploading] = useState(false);
  const [uploadingName, setUploadingName] = useState("");
  const [searchQuery, setSearchQuery] = useState("");
  const [renamingId, setRenamingId] = useState<string | null>(null);
  const [renameDraft, setRenameDraft] = useState("");
  const [confirmDeleteId, setConfirmDeleteId] = useState<string | null>(null);

  const pollRef = useRef<ReturnType<typeof setInterval> | null>(null);

  const refresh = useCallback(async () => {
    try {
      const res = await listDocuments();
      setDocuments(res.documents);
      setLoadError(null);
    } catch (err) {
      setLoadError(err instanceof ApiError ? err.message : "Failed to load documents");
    }
  }, []);

  useEffect(() => {
    (async () => {
      setLoading(true);
      await refresh();
      setLoading(false);
    })();
  }, [refresh]);

  useEffect(() => {
    const hasPending = documents.some((d) => NON_TERMINAL.has(d.status));

    if (hasPending && !pollRef.current) {
      pollRef.current = setInterval(refresh, POLL_INTERVAL_MS);
    }
    if (!hasPending && pollRef.current) {
      clearInterval(pollRef.current);
      pollRef.current = null;
    }

    return () => {
      if (pollRef.current) {
        clearInterval(pollRef.current);
        pollRef.current = null;
      }
    };
  }, [documents, refresh]);

  const upload = useCallback(
    async (file: File) => {
      setUploading(true);
      setUploadingName(file.name);
      try {
        const created = await uploadDocument(file);
        setDocuments((prev) => [
          {
            ...created,
            page_count: 0,
            summary_available: false,
            ocr_completed: false,
            embedding_generated: false,
          },
          ...prev,
        ]);
      } finally {
        setUploading(false);
        setUploadingName("");
      }
    },
    [],
  );

  const startRename = useCallback((doc: DocumentResponse) => {
    setRenamingId(doc.id);
    setRenameDraft(doc.original_name);
    setConfirmDeleteId(null);
  }, []);

  const cancelRename = useCallback(() => setRenamingId(null), []);

  const commitRename = useCallback(
    async (id: string) => {
      const name = renameDraft.trim();
      setRenamingId(null);
      if (!name) return;
      const updated = await renameDocument(id, name);
      setDocuments((prev) => prev.map((d) => (d.id === id ? updated : d)));
    },
    [renameDraft],
  );

  const startDelete = useCallback((id: string) => {
    setConfirmDeleteId(id);
    setRenamingId(null);
  }, []);

  const cancelDelete = useCallback(() => setConfirmDeleteId(null), []);

  const confirmDelete = useCallback(async (id: string) => {
    await deleteDocument(id);
    setDocuments((prev) => prev.filter((d) => d.id !== id));
    setConfirmDeleteId(null);
  }, []);

  const filteredDocuments = documents.filter(
    (d) =>
      !searchQuery.trim() ||
      d.original_name.toLowerCase().includes(searchQuery.trim().toLowerCase()),
  );

  return {
    documents: filteredDocuments,
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
  };
}
