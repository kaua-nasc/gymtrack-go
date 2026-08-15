-- Restore orphan tables and enums removed in 000010 (based on 000000/000001).

CREATE TYPE public.plan_access_request_status_enum AS ENUM ('PENDING', 'APPROVED', 'REJECTED', 'CANCELED');
CREATE TYPE public.plan_invites_status_enum AS ENUM ('PENDING', 'ACCEPTED', 'REJECTED', 'CANCELED');

CREATE TABLE public.plan_access_request (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"userId" uuid NOT NULL,
	"trainingPlanId" uuid NOT NULL,
	status public."plan_access_request_status_enum" NOT NULL,
	CONSTRAINT "PK_f03dddf7f62547e17bd18fc8df0" PRIMARY KEY (id),
	CONSTRAINT "FK_3bec98ff9eea3e72081314a4a59" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);

CREATE TABLE public.plan_invites (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"planId" uuid NOT NULL,
	"senderId" uuid NOT NULL,
	"recipientId" uuid NULL,
	"recipientEmail" varchar NOT NULL,
	status public."plan_invites_status_enum" DEFAULT 'PENDING'::plan_invites_status_enum NOT NULL,
	CONSTRAINT "PK_a948e8c5ff31526d6e5f77efbe2" PRIMARY KEY (id),
	CONSTRAINT "FK_e6b77ecc915a51b78b238e64b06" FOREIGN KEY ("planId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);

CREATE TABLE public.plan_participant (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"userId" uuid NOT NULL,
	"trainingPlanId" uuid NOT NULL,
	expiration_date timestamp NOT NULL,
	approved_at timestamp NOT NULL,
	CONSTRAINT "PK_cf83d1fc076ed13fc7447d55076" PRIMARY KEY (id),
	CONSTRAINT "FK_98e114e0b05861ac9d025e9c0af" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);

CREATE TABLE public.active_workout_sessions (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"userId" uuid NOT NULL,
	"planDayProgressId" uuid NOT NULL,
	"currentExerciseId" uuid NULL,
	"currentSetIndex" int4 DEFAULT 0 NOT NULL,
	"restStartedAt" timestamp NULL,
	"adaptiveRestDurationSeconds" int4 NULL,
	"startedAt" timestamp DEFAULT now() NOT NULL,
	"lastActiveAt" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT "PK_7c5d880eee6814a09a9ea2de6a7" PRIMARY KEY (id),
	CONSTRAINT "FK_3886d36ff4fe44eb07cc60cde74" FOREIGN KEY ("planDayProgressId") REFERENCES public.plan_day_progress(id) ON DELETE CASCADE
);

CREATE TABLE public.active_set_logs (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"sessionId" uuid NOT NULL,
	"exerciseId" uuid NOT NULL,
	"setIndex" int4 NOT NULL,
	reps int4 NOT NULL,
	weight numeric(5, 2) NOT NULL,
	rpe int4 NULL,
	CONSTRAINT "PK_cefba1f9f30da66680fe3873282" PRIMARY KEY (id),
	CONSTRAINT "FK_aa8a8d84879dac6d25ed7a3a9e6" FOREIGN KEY ("exerciseId") REFERENCES public.exercises(id) ON DELETE CASCADE,
	CONSTRAINT "FK_f6fe56d1720796ce2a7eda91689" FOREIGN KEY ("sessionId") REFERENCES public.active_workout_sessions(id) ON DELETE CASCADE
);
