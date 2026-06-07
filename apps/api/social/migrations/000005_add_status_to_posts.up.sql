CREATE TYPE public.post_status_enum AS ENUM ('PENDING', 'APPROVED', 'REJECTED');

ALTER TABLE public.posts ADD COLUMN status public.post_status_enum DEFAULT 'PENDING' NOT NULL;

UPDATE public.posts SET status = 'APPROVED';
