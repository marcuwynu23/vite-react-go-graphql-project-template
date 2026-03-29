export const theme = {
  storageKey: "theme",
  defaultMode: "system" as const,
} as const;

export type ThemeMode = "light" | "dark" | "system";
