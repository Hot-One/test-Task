CREATE TABLE phone_numbers(
    id UUID PRIMARY KEY, 
    user_id UUID REFERENCES users("id"),
    phone varchar,
    description varchar,
    is_fax boolean
);