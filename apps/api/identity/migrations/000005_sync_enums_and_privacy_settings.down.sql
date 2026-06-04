ALTER TABLE public.metric_goals ALTER COLUMN "type" TYPE varchar;
DROP TYPE public.metric_goals_type_enum;

-- Note: Removing values from an ENUM is not supported in Postgres.
