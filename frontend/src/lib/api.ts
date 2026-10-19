import type {
  DocumentFileResponse,
  DocumentResponse,
  ListChunksResponse,
  ListDocumentsResponse,
  UploadDocumentResponse,
} from "./types";

export const API_BASE = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:9090";

export class ApiError extends Error {
  status: number;

  constructor(status: number, message: string) {
    super(message);
    this.name = "ApiError";
    this.status = status;
  }
}

export async function buildApiError(res: Response): Promise<ApiError> {
  let message = `Request failed with status ${res.status}`;
  try {
    const body = (await res.json()) as { error?: string };
    if (body.error) message = body.error;
  } catch {
    // response had no JSON body
  }
  return new ApiError(res.status, message);
}

async function handleResponse<T>(res: Response): Promise<T> {
  if (!res.ok) {
    throw await buildApiError(res);
  }
  return res.json() as Promise<T>;
}

export async function listDocuments(): Promise<ListDocumentsResponse> {
  const res = await fetch(`${API_BASE}/api/v1/documents`);
  return handleResponse<ListDocumentsResponse>(res);
}

export async function getDocument(id: string): Promise<DocumentResponse> {
  const res = await fetch(`${API_BASE}/api/v1/documents/${id}`);
  return handleResponse<DocumentResponse>(res);
}

export async function uploadDocument(
  file: File,
): Promise<UploadDocumentResponse> {
  const formData = new FormData();
  formData.append("file", file);
  const res = await fetch(`${API_BASE}/api/v1/documents`, {
    method: "POST",
    body: formData,
  });
  return handleResponse<UploadDocumentResponse>(res);
}

export async function renameDocument(
  id: string,
  name: string,
): Promise<DocumentResponse> {
  const res = await fetch(`${API_BASE}/api/v1/documents/${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name }),
  });
  return handleResponse<DocumentResponse>(res);
}

export async function deleteDocument(
  id: string,
): Promise<{ message: string }> {
  const res = await fetch(`${API_BASE}/api/v1/documents/${id}`, {
    method: "DELETE",
  });
  return handleResponse<{ message: string }>(res);
}

export async function getDocumentFile(
  id: string,
): Promise<DocumentFileResponse> {
  const res = await fetch(`${API_BASE}/api/v1/documents/${id}/file`);
  return handleResponse<DocumentFileResponse>(res);
}

export async function getDocumentChunks(
  id: string,
): Promise<ListChunksResponse> {
  const res = await fetch(`${API_BASE}/api/v1/documents/${id}/chunks`);
  return handleResponse<ListChunksResponse>(res);
}
