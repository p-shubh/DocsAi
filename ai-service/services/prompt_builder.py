def build_prompt(
    question: str,
    context: list[dict],
) -> str:

    context_blocks = []

    for chunk in context:
        context_blocks.append(
            f"[Source: {chunk['document_name']}, "
            f"Page {chunk['page']}, "
            f"Section: {chunk['section'] or 'N/A'}]\n"
            f"{chunk['text']}"
        )

    context_text = "\n\n".join(context_blocks)

    return (
        "You are DocsAI.\n"
        "Answer ONLY using the provided context.\n"
        "If the answer cannot be found, say you don't know.\n\n"
        f"Context:\n{context_text}\n\n"
        f"Question:\n{question}"
    )
