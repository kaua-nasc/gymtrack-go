ALTER TABLE public.days DROP CONSTRAINT IF EXISTS "UQ_days_plan_name";
CREATE UNIQUE INDEX "UQ_days_plan_name_active" ON public.days ("trainingPlanId", name) WHERE "deletedAt" IS NULL;
