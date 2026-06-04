-- Add CANCELED to plan_day_progress_status_enum
ALTER TYPE public.plan_day_progress_status_enum ADD VALUE 'CANCELED';

-- Update existing data to use the new spelling
UPDATE public.plan_day_progress SET status = 'CANCELED' WHERE status = 'CANCELLED';

-- Ensure plan_invites_status_enum and plan_access_request_status_enum have CANCELED (with one L)
-- If they were created with CANCELLED, we add the correct spelling
ALTER TYPE public.plan_invites_status_enum ADD VALUE 'CANCELED';
ALTER TYPE public.plan_access_request_status_enum ADD VALUE 'CANCELED';
