CREATE TABLE post_audit_logs (
    id UUID PRIMARY KEY,
    "postId" UUID NOT NULL REFERENCES posts(id),
    "adminId" UUID NOT NULL,
    "previousStatus" post_status_enum NOT NULL,
    "newStatus" post_status_enum NOT NULL,
    reason TEXT,
    "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_post_audit_logs_created_at ON post_audit_logs ("createdAt" DESC);
CREATE INDEX idx_post_audit_logs_new_status ON post_audit_logs ("newStatus");
