-- "maxSubscriptions" was nullable (NULL meant "unlimited"). Now 0 means "unlimited".

UPDATE public.training_plans SET "maxSubscriptions" = 0 WHERE "maxSubscriptions" IS NULL;

ALTER TABLE public.training_plans ALTER COLUMN "maxSubscriptions" SET DEFAULT 0;
ALTER TABLE public.training_plans ALTER COLUMN "maxSubscriptions" SET NOT NULL;