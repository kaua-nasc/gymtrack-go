-- Restore the nullable "maxSubscriptions" column where 0 meant "unlimited" via NULL.

ALTER TABLE public.training_plans ALTER COLUMN "maxSubscriptions" DROP DEFAULT;
ALTER TABLE public.training_plans ALTER COLUMN "maxSubscriptions" DROP NOT NULL;

UPDATE public.training_plans SET "maxSubscriptions" = NULL WHERE "maxSubscriptions" = 0;