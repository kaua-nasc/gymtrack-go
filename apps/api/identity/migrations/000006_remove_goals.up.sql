-- Remove metric_goals table
DROP TABLE IF EXISTS public.metric_goals;

-- Remove shareMetricGoals column from user_privacy_settings
ALTER TABLE public.user_privacy_settings DROP COLUMN IF EXISTS "shareMetricGoals";

-- Remove enum types
DROP TYPE IF EXISTS public.metric_goals_status_enum;
DROP TYPE IF EXISTS public.metric_goals_type_enum;
