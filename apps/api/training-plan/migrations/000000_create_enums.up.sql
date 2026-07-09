-- Enums for public.training_plans
CREATE TYPE public.training_plans_type_enum AS ENUM ('HYPERTROPHY', 'STRENGTH', 'MIXED');
CREATE TYPE public.training_plans_visibility_enum AS ENUM ('PUBLIC', 'PROTECTED', 'PRIVATE');
CREATE TYPE public.training_plans_level_enum AS ENUM ('BEGINNER', 'INTERMEDIATE', 'ADVANCED');

-- Enums for public.exercises
CREATE TYPE public.exercises_type_enum AS ENUM ('WARMUP', 'RECOGNITION', 'WORK', 'CARDIO');

-- Enums for public.plan_access_request
CREATE TYPE public.plan_access_request_status_enum AS ENUM ('PENDING', 'APPROVED', 'REJECTED', 'CANCELED');

-- Enums for public.plan_invites
CREATE TYPE public.plan_invites_status_enum AS ENUM ('PENDING', 'ACCEPTED', 'REJECTED', 'CANCELED');

-- Enums for public.plan_subscription
CREATE TYPE public.plan_subscription_status_enum AS ENUM ('NOT_STARTED', 'IN_PROGRESS', 'COMPLETED', 'CANCELED');
CREATE TYPE public.plan_subscription_type_enum AS ENUM ('TOTAL_ACCESS', 'PARTIAL_ACCESS', 'PRIVATE');

-- Enums for public.plan_day_progress
CREATE TYPE public.plan_day_progress_status_enum AS ENUM ('IN_PROGRESS', 'COMPLETED', 'CANCELED');
