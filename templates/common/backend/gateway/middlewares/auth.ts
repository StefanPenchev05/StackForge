import { Request, Response, NextFunction } from "express";

export default function authMiddleware(req: Request, res: Response, next: NextFunction) {
  const token = req.headers.authorization;
  if (!token || token !== "Bearer mysecrettoken") {
    res.status(401).json({ message: "Unauthorized" });
    return;
  }
  next();
}
