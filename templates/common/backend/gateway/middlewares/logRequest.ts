import { Request, Response, NextFunction } from "express";

export default function logRequest(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const timestamp = new Date().toISOString();
  const method = req.method;
  const url = req.originalUrl;
  const ip =
    req.ip || req.headers["x-forwarded-for"] || req.connection.remoteAddress;

  console.log(`[${timestamp}] ${method} ${url} - IP: ${ip}`);

  next();
}
