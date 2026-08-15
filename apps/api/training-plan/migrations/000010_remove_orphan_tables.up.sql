-- Remove orphan tables and enums left from legacy/decoupled designs:
-- active_workout_sessions and active_set_logs (never implemented),
-- plan_access_request, plan_invites and plan_participant (no feature uses them).

DROP TABLE IF EXISTS public.active_set_logs;
DROP TABLE IF EXISTS public.active_workout_sessions;
DROP TABLE IF EXISTS public.plan_participant;
DROP TABLE IF EXISTS public.plan_invites;
DROP TABLE IF EXISTS public.plan_access_request;

DROP TYPE IF EXISTS public.plan_invites_status_enum;
DROP TYPE IF EXISTS public.plan_access_request_status_enum;
