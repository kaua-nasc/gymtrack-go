-- Create enum type for post entity type
CREATE TYPE public.post_entity_type_enum AS ENUM ('TRAINING_PLAN');

-- Alter table to use the new enum type
ALTER TABLE public.posts 
ALTER COLUMN "entityType" TYPE public.post_entity_type_enum 
USING "entityType"::public.post_entity_type_enum;