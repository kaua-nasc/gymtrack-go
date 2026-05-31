-- Adiciona novos tipos ao ENUM
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'BODY_FAT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'WATER_PERCENTAGE';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'MUSCLE_MASS';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'BONE_MASS';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'NECK';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'SHOULDERS';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'BICEP_LEFT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'BICEP_RIGHT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'FOREARM_LEFT';
ALTER TYPE public.body_measurements_type_enum ADD VALUE 'FOREARM_RIGHT';

-- Migra dados existentes de Braço para Bíceps
-- Nota: Usamos COMMIT/COMMIT para garantir que os novos valores do ENUM estejam disponíveis para o UPDATE se rodado em blocos separados,
-- mas em migrações SQL puras o Postgres geralmente exige que o ALTER TYPE ADD VALUE rode fora de uma transação ou que a transação seja concluída.
-- No entanto, a maioria dos drivers de migração lida com isso.

UPDATE public.body_measurements SET "type" = 'BICEP_LEFT' WHERE "type" = 'ARM_LEFT';
UPDATE public.body_measurements SET "type" = 'BICEP_RIGHT' WHERE "type" = 'ARM_RIGHT';
