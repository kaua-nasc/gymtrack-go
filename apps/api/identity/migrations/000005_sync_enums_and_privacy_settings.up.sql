-- Expand body_measurements_type_enum
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'THIGH_LEFT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'THIGH_RIGHT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'CALF_LEFT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'CALF_RIGHT';

-- Create metric_goals_type_enum
CREATE TYPE public.metric_goals_type_enum AS ENUM ('WEIGHT', 'BODY_FAT', 'MUSCLE_MASS', 'BONE_MASS', 'CHEST', 'WAIST', 'HIPS');

-- Convert metric_goals.type to enum
ALTER TABLE public.metric_goals 
ALTER COLUMN "type" TYPE public.metric_goals_type_enum 
USING "type"::public.metric_goals_type_enum;
