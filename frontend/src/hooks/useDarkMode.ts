"use client";

import { useCallback, useEffect, useState } from "react";

const STORAGE_KEY = "docsai-theme";

export function useDarkMode() {
  const [dark, setDark] = useState(false);

  useEffect(() => {
    const stored = window.localStorage.getItem(STORAGE_KEY);
    const initial = stored
      ? stored === "dark"
      : window.matchMedia("(prefers-color-scheme: dark)").matches;
    // One-time read of a browser-only preference source (localStorage /
    // matchMedia) on mount — there is no server-known value to derive this
    // from, so it can't be computed via a lazy useState initializer without
    // causing an SSR/client hydration mismatch instead.
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setDark(initial);
    document.documentElement.classList.toggle("dark", initial);
  }, []);

  const toggle = useCallback(() => {
    setDark((prev) => {
      const next = !prev;
      document.documentElement.classList.toggle("dark", next);
      window.localStorage.setItem(STORAGE_KEY, next ? "dark" : "light");
      return next;
    });
  }, []);

  return { dark, toggle };
}
