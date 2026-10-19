"use client";

import { useRef, useState } from "react";

export function UploadDropzone({
  uploading,
  uploadingName,
  onUpload,
}: {
  uploading: boolean;
  uploadingName: string;
  onUpload: (file: File) => void;
}) {
  const [dragActive, setDragActive] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  return (
    <div
      role="button"
      tabIndex={0}
      onClick={() => inputRef.current?.click()}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") inputRef.current?.click();
      }}
      onDragOver={(e) => {
        e.preventDefault();
        setDragActive(true);
      }}
      onDragLeave={(e) => {
        e.preventDefault();
        setDragActive(false);
      }}
      onDrop={(e) => {
        e.preventDefault();
        setDragActive(false);
        const file = e.dataTransfer.files?.[0];
        if (file) onUpload(file);
      }}
      className={`mb-[22px] cursor-pointer rounded-xl border-[1.5px] border-dashed p-7 text-center transition-colors ${
        dragActive ? "border-accent bg-accent-soft" : "border-border bg-surface"
      }`}
    >
      <input
        ref={inputRef}
        type="file"
        className="hidden"
        onChange={(e) => {
          const file = e.target.files?.[0];
          if (file) onUpload(file);
          e.target.value = "";
        }}
      />
      <div className="mx-auto flex h-[34px] w-[34px] items-center justify-center rounded-[9px] bg-accent-soft text-base font-bold text-accent">
        ↑
      </div>
      <div className="mt-2.5 text-sm font-semibold text-ink">
        {uploading ? `Uploading ${uploadingName}…` : "Drag & drop a file here"}
      </div>
      <div className="mt-1 text-[12.5px] text-ink-faint">
        or click to browse — one file at a time
      </div>
    </div>
  );
}
