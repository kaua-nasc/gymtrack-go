-- Note: Updating data back to CANCELLED might be necessary if we wanted a full rollback, 
-- but ENUM values cannot be removed.
UPDATE public.plan_day_progress SET status = 'CANCELLED' WHERE status = 'CANCELED';
