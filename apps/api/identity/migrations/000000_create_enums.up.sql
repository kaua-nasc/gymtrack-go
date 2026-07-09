-- Create uuid-ossp extension (required for uuid_generate_v4 in existing migrations)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums for public.users
CREATE TYPE public.users_type_enum AS ENUM ('PERSONAL_TRAINER', 'CLIENT');
CREATE TYPE public.users_weightunit_enum AS ENUM ('kg', 'lb');
CREATE TYPE public.users_heightunit_enum AS ENUM ('cm', 'in');

-- Enums for public.body_measurements
CREATE TYPE public.body_measurements_type_enum AS ENUM ('CHEST', 'WAIST', 'HIPS', 'ARM_LEFT', 'ARM_RIGHT');

-- Enums for public.metric_goals
CREATE TYPE public.metric_goals_status_enum AS ENUM ('ACTIVE');
