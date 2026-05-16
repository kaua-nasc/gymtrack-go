-- Create posts table
CREATE TABLE public.posts (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"authorId" uuid NOT NULL,
	"content" text NOT NULL,
	"entityId" uuid NOT NULL,
	"entityType" varchar NOT NULL,
	CONSTRAINT "PK_posts" PRIMARY KEY (id)
);

-- Create post_likes table
CREATE TABLE public.post_likes (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"userId" uuid NOT NULL,
	"postId" uuid NOT NULL,
	CONSTRAINT "PK_post_likes" PRIMARY KEY (id),
	CONSTRAINT "FK_post_likes_postId" FOREIGN KEY ("postId") REFERENCES public.posts(id) ON DELETE CASCADE
);

-- Create post_comments table
CREATE TABLE public.post_comments (
	id uuid NOT NULL,
	"createdAt" timestamp DEFAULT now() NOT NULL,
	"updatedAt" timestamp DEFAULT now() NOT NULL,
	"deletedAt" timestamp NULL,
	"content" text NOT NULL,
	"authorId" uuid NOT NULL,
	"postId" uuid NOT NULL,
	CONSTRAINT "PK_post_comments" PRIMARY KEY (id),
	CONSTRAINT "FK_post_comments_postId" FOREIGN KEY ("postId") REFERENCES public.posts(id) ON DELETE CASCADE
);
