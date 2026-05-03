-- public.training_plans definition

-- Drop table

-- DROP TABLE public.training_plans;

CREATE TABLE public.training_plans (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"name" varchar NOT NULL,
	"authorId" uuid NOT NULL,
	"timeInDays" int4 NOT NULL,
	"type" public."training_plans_type_enum" NOT NULL,
	observation text NULL,
	pathology text NULL,
	visibility public."training_plans_visibility_enum" NOT NULL,
	"level" public."training_plans_level_enum" NOT NULL,
	"maxSubscriptions" int4 NULL,
	"imageUrl" text NULL,
	description text NULL,
	CONSTRAINT "PK_246975cb895b51662b90515a390" PRIMARY KEY (id)
);


-- public.days definition

-- Drop table

-- DROP TABLE public.days;

CREATE TABLE public.days (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"name" varchar NOT NULL,
	"trainingPlanId" uuid NOT NULL,
	CONSTRAINT "PK_c2c66eb46534bea34ba48cc4d7f" PRIMARY KEY (id),
	CONSTRAINT "UQ_18a49040122a90a7f959508f9b8" UNIQUE (name),
	CONSTRAINT "FK_4c4841535803ef06571a77782fc" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);


-- public.exercises definition

-- Drop table

-- DROP TABLE public.exercises;

CREATE TABLE public.exercises (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"name" varchar NOT NULL,
	"dayId" uuid NOT NULL,
	"type" public."exercises_type_enum" NOT NULL,
	"setsNumber" int4 NOT NULL,
	"repsNumber" int4 NOT NULL,
	description text NULL,
	observation text NULL,
	CONSTRAINT "PK_c4c46f5fa89a58ba7c2d894e3c3" PRIMARY KEY (id),
	CONSTRAINT "UQ_a521b5cac5648eedc036e17d1bd" UNIQUE (name),
	CONSTRAINT "FK_85531791853605820c4f905ec7a" FOREIGN KEY ("dayId") REFERENCES public.days(id) ON DELETE CASCADE
);



-- public.plan_access_request definition

-- Drop table

-- DROP TABLE public.plan_access_request;

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


-- public.plan_invites definition

-- Drop table

-- DROP TABLE public.plan_invites;

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


-- public.plan_participant definition

-- Drop table

-- DROP TABLE public.plan_participant;

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


-- public.plan_subscription definition

-- Drop table

-- DROP TABLE public.plan_subscription;

CREATE TABLE public.plan_subscription (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"trainingPlanId" uuid NOT NULL,
	"userId" uuid NOT NULL,
	status public."plan_subscription_status_enum" DEFAULT 'NOT_STARTED'::plan_subscription_status_enum NOT NULL,
	"type" public."plan_subscription_type_enum" NOT NULL,
	CONSTRAINT "PK_537e7826b55596d075de1bde618" PRIMARY KEY (id),
	CONSTRAINT "FK_7c57629907dcde0a2029c5dc68a" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);


-- public.plan_subscription_privacy_settings definition

-- Drop table

-- DROP TABLE public.plan_subscription_privacy_settings;

CREATE TABLE public.plan_subscription_privacy_settings (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"shareProgress" bool DEFAULT true NOT NULL,
	"sharePersonalMetrics" bool DEFAULT false NOT NULL,
	"planSubscriptionId" uuid NULL,
	CONSTRAINT "PK_ed7f77b1c435440568ecff621e4" PRIMARY KEY (id),
	CONSTRAINT "REL_f0fe6f036ae32d5bfe1af9dc13" UNIQUE ("planSubscriptionId"),
	CONSTRAINT "FK_f0fe6f036ae32d5bfe1af9dc133" FOREIGN KEY ("planSubscriptionId") REFERENCES public.plan_subscription(id) ON DELETE CASCADE
);


-- public.training_plan_comments definition

-- Drop table

-- DROP TABLE public.training_plan_comments;

CREATE TABLE public.training_plan_comments (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"content" text NOT NULL,
	"authorId" uuid NOT NULL,
	"trainingPlanId" uuid NOT NULL,
	CONSTRAINT "PK_f9ae0ded8ad0835f8598e41963a" PRIMARY KEY (id),
	CONSTRAINT "FK_9e1eb349d94fb4593ff8bace996" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);


-- public.training_plan_feedbacks definition

-- Drop table

-- DROP TABLE public.training_plan_feedbacks;

CREATE TABLE public.training_plan_feedbacks (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"trainingPlanId" uuid NOT NULL,
	"userId" uuid NOT NULL,
	rating numeric NOT NULL,
	message text NULL,
	CONSTRAINT "PK_220c6f963204b4ee178c6e3bd16" PRIMARY KEY (id),
	CONSTRAINT "FK_f1cfb6d6ba2120341efcac05041" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);


-- public.training_plan_likes definition

-- Drop table

-- DROP TABLE public.training_plan_likes;

CREATE TABLE public.training_plan_likes (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"likedBy" uuid NOT NULL,
	"trainingPlanId" uuid NOT NULL,
	CONSTRAINT "PK_1034857b4d99c260c6a4d9218e5" PRIMARY KEY (id),
	CONSTRAINT "FK_2c85f41f867f9d0703b165425e9" FOREIGN KEY ("trainingPlanId") REFERENCES public.training_plans(id) ON DELETE CASCADE
);

-- public.exercise_logs definition

-- Drop table

-- DROP TABLE public.exercise_logs;

CREATE TABLE public.exercise_logs (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"userId" uuid NOT NULL,
	"exerciseId" uuid NOT NULL,
	reps text NOT NULL,
	weight text NOT NULL,
	notes text NULL,
	CONSTRAINT "PK_32076bf978e4169be16e25bf8dc" PRIMARY KEY (id),
	CONSTRAINT "FK_1c25082e5788b58e4a91407e34d" FOREIGN KEY ("exerciseId") REFERENCES public.exercises(id) ON DELETE CASCADE
);


-- public.plan_day_progress definition

-- Drop table

-- DROP TABLE public.plan_day_progress;

CREATE TABLE public.plan_day_progress (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"planSubscriptionId" uuid NOT NULL,
	"dayId" uuid NOT NULL,
	status public."plan_day_progress_status_enum" DEFAULT 'IN_PROGRESS'::plan_day_progress_status_enum NOT NULL,
	CONSTRAINT "PK_6a1d175ad26bc8e79d4253a447a" PRIMARY KEY (id),
	CONSTRAINT "FK_01e9018d0b5401f2d6bc02fca51" FOREIGN KEY ("planSubscriptionId") REFERENCES public.plan_subscription(id) ON DELETE CASCADE,
	CONSTRAINT "FK_4ad012ce3c8a11b973cb7d421d6" FOREIGN KEY ("dayId") REFERENCES public.days(id) ON DELETE CASCADE
);


-- public.active_workout_sessions definition

-- Drop table

-- DROP TABLE public.active_workout_sessions;

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


-- public.active_set_logs definition

-- Drop table

-- DROP TABLE public.active_set_logs;

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