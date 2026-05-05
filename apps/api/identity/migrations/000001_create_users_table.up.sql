-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"firstName" varchar NOT NULL,
	"lastName" varchar NOT NULL,
	email varchar NOT NULL,
	bio text NULL,
	"profilePictureUrl" text NULL,
	"password" varchar NOT NULL,
	"type" public."users_type_enum" DEFAULT 'CLIENT'::users_type_enum NOT NULL,
	height numeric(5, 2) NULL,
	"currentWeight" numeric(5, 2) NULL,
	"weightUnit" public."users_weightunit_enum" DEFAULT 'kg'::users_weightunit_enum NOT NULL,
	"heightUnit" public."users_heightunit_enum" DEFAULT 'cm'::users_heightunit_enum NOT NULL,
	"trainerInviteCode" varchar(50) NULL,
	cref varchar(20) NULL,
	"isVerified" bool DEFAULT false NOT NULL,
	CONSTRAINT "PK_a3ffb1c0c8416b9fc6f907b7433" PRIMARY KEY (id),
	CONSTRAINT "UQ_52ad64cb712ed325dff463f2f13" UNIQUE ("trainerInviteCode"),
	CONSTRAINT "UQ_584f11b3c6df319615b141165a2" UNIQUE (cref),
	CONSTRAINT "UQ_97672ac88f789774dd47f7c8be3" UNIQUE (email)
);


-- public.body_measurements definition

-- Drop table

-- DROP TABLE public.body_measurements;

CREATE TABLE public.body_measurements (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"type" public."body_measurements_type_enum" NOT NULL,
	value numeric(6, 2) NOT NULL,
	"measuredAt" timestamp DEFAULT now() NOT NULL,
	"userId" uuid NOT NULL,
	"trainerNote" text NULL,
	"trainerNoteAt" timestamp NULL,
	CONSTRAINT "PK_474282e620ea0cd4fe5d8cbce0f" PRIMARY KEY (id),
	CONSTRAINT "FK_e52cb7540788161ebe6a1ac2113" FOREIGN KEY ("userId") REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.metric_goals definition

-- Drop table

-- DROP TABLE public.metric_goals;

CREATE TABLE public.metric_goals (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"type" varchar NOT NULL,
	"startingValue" numeric(6, 2) NOT NULL,
	"targetValue" numeric(6, 2) NOT NULL,
	deadline timestamp NULL,
	"achievedAt" timestamp NULL,
	status public."metric_goals_status_enum" DEFAULT 'ACTIVE'::metric_goals_status_enum NOT NULL,
	"userId" uuid NOT NULL,
	CONSTRAINT "PK_4b91e2d0945f2434af811ea86df" PRIMARY KEY (id),
	CONSTRAINT "FK_920692c5277515acb1536e08f64" FOREIGN KEY ("userId") REFERENCES public.users(id) ON DELETE CASCADE
);

-- public.trainer_student_relationships definition

-- Drop table

-- DROP TABLE public.trainer_student_relationships;

CREATE TABLE public.trainer_student_relationships (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"trainerId" uuid NOT NULL,
	"studentId" uuid NOT NULL,
	"linkedAt" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT "PK_7bcad11b188ec201c383fb768d6" PRIMARY KEY (id),
	CONSTRAINT "UQ_6466809dbe84657bfe692add1ce" UNIQUE ("studentId"),
	CONSTRAINT "FK_6466809dbe84657bfe692add1ce" FOREIGN KEY ("studentId") REFERENCES public.users(id),
	CONSTRAINT "FK_abb1b731e39687111e7feb0288e" FOREIGN KEY ("trainerId") REFERENCES public.users(id)
);


-- public.user_follows definition

-- Drop table

-- DROP TABLE public.user_follows;

CREATE TABLE public.user_follows (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"followerId" uuid NOT NULL,
	"followingId" uuid NOT NULL,
	CONSTRAINT "PK_da8e8793113adf3015952880966" PRIMARY KEY (id),
	CONSTRAINT "FK_6300484b604263eaae8a6aab88d" FOREIGN KEY ("followerId") REFERENCES public.users(id) ON DELETE CASCADE,
	CONSTRAINT "FK_7c6c27f12c4e972eab4b3aaccbf" FOREIGN KEY ("followingId") REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.user_privacy_settings definition

-- Drop table

-- DROP TABLE public.user_privacy_settings;

CREATE TABLE public.user_privacy_settings (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"shareName" bool DEFAULT true NOT NULL,
	"shareEmail" bool DEFAULT true NOT NULL,
	"shareTrainingProgress" bool DEFAULT false NOT NULL,
	"sharePastDataWithTrainer" bool DEFAULT false NOT NULL,
	"userId" uuid NULL,
	CONSTRAINT "PK_95fc563e79fedf0b241c0360be0" PRIMARY KEY (id),
	CONSTRAINT "REL_6c6227fa8fb10ca8cc23ce6939" UNIQUE ("userId"),
	CONSTRAINT "FK_6c6227fa8fb10ca8cc23ce6939b" FOREIGN KEY ("userId") REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.weight_logs definition

-- Drop table

-- DROP TABLE public.weight_logs;

CREATE TABLE public.weight_logs (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	weight numeric(5, 2) NOT NULL,
	"measuredAt" timestamp DEFAULT now() NOT NULL,
	"userId" uuid NOT NULL,
	"trainerNote" text NULL,
	"trainerNoteAt" timestamp NULL,
	CONSTRAINT "PK_96c8f4d341846b34fef50cf4576" PRIMARY KEY (id),
	CONSTRAINT "FK_5d83c1656fef64dfbc54ab7ce37" FOREIGN KEY ("userId") REFERENCES public.users(id) ON DELETE CASCADE
);
