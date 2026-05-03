-- Initial Migration for Training Plan Module

CREATE TABLE IF NOT EXISTS training_plans (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author_id UUID NOT NULL,
    time_in_days INTEGER NOT NULL,
    type VARCHAR(50) NOT NULL,
    visibility VARCHAR(50) NOT NULL,
    level VARCHAR(50) NOT NULL,
    observation TEXT,
    pathology TEXT,
    max_subscriptions INTEGER,
    image_url TEXT,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS days (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exercises (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    day_id UUID NOT NULL REFERENCES days(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    sets_number INTEGER NOT NULL,
    reps_number INTEGER NOT NULL,
    description TEXT,
    observation TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_plan_likes (
    id VARCHAR(255) PRIMARY KEY, -- Using composite id "plan:user"
    liked_by UUID NOT NULL,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(liked_by, training_plan_id)
);

CREATE TABLE IF NOT EXISTS training_plan_comments (
    id VARCHAR(255) PRIMARY KEY,
    content TEXT NOT NULL,
    author_id UUID NOT NULL,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS plan_subscriptions (
    id VARCHAR(255) PRIMARY KEY, -- Using composite id "plan:user"
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, training_plan_id)
);

CREATE TABLE IF NOT EXISTS plan_day_progress (
    id VARCHAR(255) PRIMARY KEY,
    day_id UUID NOT NULL REFERENCES days(id) ON DELETE CASCADE,
    plan_subscription_id VARCHAR(255) NOT NULL REFERENCES plan_subscriptions(id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS private_participants (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    expiration_date TIMESTAMP WITH TIME ZONE,
    approved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, training_plan_id)
);

CREATE TABLE IF NOT EXISTS plan_access_requests (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS plan_invites (
    id UUID PRIMARY KEY,
    plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL,
    recipient_id UUID,
    recipient_email VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_plan_feedbacks (
    id VARCHAR(255) PRIMARY KEY,
    training_plan_id UUID NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    rating DECIMAL(3,2) NOT NULL,
    message TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exercise_logs (
    id VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL,
    exercise_id UUID NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    reps INTEGER[] NOT NULL,
    weight DECIMAL[] NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
