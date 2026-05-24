DROP INDEX IF EXISTS "UQ_trainer_student_active_student";
ALTER TABLE public.trainer_student_relationships ADD CONSTRAINT "UQ_6466809dbe84657bfe692add1ce" UNIQUE ("studentId");
