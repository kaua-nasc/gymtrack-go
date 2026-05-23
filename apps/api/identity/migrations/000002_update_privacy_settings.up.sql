ALTER TABLE public.user_privacy_settings DROP COLUMN "shareName";
ALTER TABLE public.user_privacy_settings ADD COLUMN "shareBodyMeasurements" bool DEFAULT false NOT NULL;
ALTER TABLE public.user_privacy_settings ADD COLUMN "shareWeightLogs" bool DEFAULT false NOT NULL;
ALTER TABLE public.user_privacy_settings ADD COLUMN "shareMetricGoals" bool DEFAULT false NOT NULL;
ALTER TABLE public.user_privacy_settings ADD COLUMN "allowTrainerNotes" bool DEFAULT true NOT NULL;
