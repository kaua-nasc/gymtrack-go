ALTER TABLE public.user_privacy_settings DROP COLUMN "shareBodyMeasurements";
ALTER TABLE public.user_privacy_settings DROP COLUMN "shareWeightLogs";
ALTER TABLE public.user_privacy_settings DROP COLUMN "shareMetricGoals";
ALTER TABLE public.user_privacy_settings DROP COLUMN "allowTrainerNotes";
ALTER TABLE public.user_privacy_settings ADD COLUMN "shareName" bool DEFAULT true NOT NULL;
