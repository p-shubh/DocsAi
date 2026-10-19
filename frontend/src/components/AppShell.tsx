"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useDarkMode } from "@/hooks/useDarkMode";

const NAV_ITEMS = [
  { href: "/documents", label: "Documents" },
  { href: "/chat", label: "Chat" },
];

export function AppShell({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const { dark, toggle } = useDarkMode();

  return (
    <div className="flex min-h-screen w-full flex-col bg-bg text-ink">
      <div className="flex h-[58px] shrink-0 items-center gap-[14px] border-b border-border bg-surface px-7">
        <div className="flex items-center gap-2.5">
          <div className="flex h-6 w-6 items-center justify-center rounded-[7px] bg-accent text-xs font-bold text-white">
            D
          </div>
          <span className="text-[15px] font-bold tracking-tight">DocsAI</span>
        </div>

        <nav className="ml-4 flex flex-1 gap-1">
          {NAV_ITEMS.map((item) => {
            const active = pathname?.startsWith(item.href);
            return (
              <Link
                key={item.href}
                href={item.href}
                className={`rounded-lg px-3.5 py-2 text-[13.5px] font-medium ${
                  active
                    ? "bg-accent-soft text-accent font-semibold"
                    : "text-ink-muted hover:text-ink"
                }`}
              >
                {item.label}
              </Link>
            );
          })}
        </nav>

        <button
          type="button"
          title="Toggle theme"
          onClick={toggle}
          className="rounded p-1.5 text-sm text-ink-faint hover:text-ink"
        >
          {dark ? "☀" : "☾"}
        </button>

        <div className="flex h-7 w-7 shrink-0 items-center justify-center rounded-full bg-hover text-xs font-semibold text-ink-muted">
          U
        </div>
      </div>

      <div className="flex flex-1 min-h-0 flex-col">{children}</div>
    </div>
  );
}
