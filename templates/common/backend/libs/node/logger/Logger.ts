export type LogLevel = "info" | "warn" | "error" | "debug";

export class Logger {
  private static format(
    level: LogLevel,
    message: string,
    context?: string
  ): string {
    const timestamp = new Date().toISOString();
    const ctx = context ? `[${context}] ` : "";
    return `[${timestamp}] [${level.toUpperCase()}] ${ctx}${message}`;
  }

  static info(message: string, context?: string) {
    console.log(this.format("info", message, context));
  }

  static warn(message: string, context?: string) {
    console.warn(this.format("warn", message, context));
  }

  static error(message: string, context?: string, error?: any) {
    console.error(this.format("error", message, context));
    if (error) console.error(error);
  }

  static debug(message: string, context?: string) {
    console.debug(this.format("debug", message, context));
  }
}
