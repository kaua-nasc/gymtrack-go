-- Restore enum types
CREATE TYPE public.metric_goals_status_enum AS ENUM ('ACTIVE', 'ACHIEVED', 'ABANDONED');
CREATE TYPE public.metric_goals_type_enum AS ENUM ('WEIGHT', 'BODY_FAT', 'MUSCLE_MASS', 'BONE_MASS', 'CHEST', 'WAIST', 'HIPS');

-- Restore shareMetricGoals column
ALTER TABLE public.user_privacy_settings ADD COLUMN "shareMetricGoals" bool DEFAULT false NOT NULL;

-- Restore metric_goals table
CREATE TABLE public.metric_goals (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"type" public.metric_goals_type_enum NOT NULL,
	"startingValue" numeric(6, 2) NOT NULL,
	"targetValue" numeric(6, 2) NOT NULL,
	deadline timestamp NULL,
	"achievedAt" timestamp NULL,
	status public.metric_goals_status_enum DEFAULT 'ACTIVE'::metric_goals_status_enum NOT NULL,
	"userId" uuid NOT NULL,
	CONSTRAINT "PK_4b91e2d0945f2434af811ea86df" PRIMARY KEY (id),
	CONSTRAINT "FK_920692c5277515acb1536e08f64" FOREIGN KEY ("userId") REFERENCES public.users(id) ON DELETE CASCADE
);
