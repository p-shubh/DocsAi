export type DocumentStatus =
  | "uploaded"
  | "queued"
  | "processing"
  | "completed"
  | "failed";

export interface DocumentResponse {
  id: string;
  original_name: string;
  size: number;
  mime_type: string;
  status: DocumentStatus;
  page_count: number;
  summary_available: boolean;
  ocr_completed: boolean;
  embedding_generated: boolean;
  uploaded_at: string;
}

export interface UploadDocumentResponse {
  id: string;
  original_name: string;
  size: number;
  mime_type: string;
  status: DocumentStatus;
  uploaded_at: string;
}

export interface ListDocumentsResponse {
  documents: DocumentResponse[];
  total: number;
  page: number;
  limit: number;
}

export interface SourceCitation {
  document_name: string;
  page: number;
  section: string;
  chunk_id: string;
}

export interface AskQuestionResponse {
  answer: string;
  model: string;
  sources: SourceCitation[];
}

export interface DocumentFileResponse {
  file_url: string;
  mime_type: string;
  original_name: string;
}

export interface DocumentChunk {
  chunk_index: number;
  page: number;
  section: string;
  text: string;
}

export interface ListChunksResponse {
  chunks: DocumentChunk[];
}
