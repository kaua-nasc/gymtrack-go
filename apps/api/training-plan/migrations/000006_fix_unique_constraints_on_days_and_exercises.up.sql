ALTER TABLE public.days DROP CONSTRAINT "UQ_18a49040122a90a7f959508f9b8";
ALTER TABLE public.days ADD CONSTRAINT "UQ_days_plan_name" UNIQUE ("trainingPlanId", name);

ALTER TABLE public.exercises DROP CONSTRAINT "UQ_a521b5cac5648eedc036e17d1bd";
ALTER TABLE public.exercises ADD CONSTRAINT "UQ_exercises_day_name" UNIQUE ("dayId", name);
