   CREATE TYPE user_role AS ENUM ('game_master','player');
   CREATE TYPE character_class AS ENUM ('warrior','mage'); --добавить клирика, вора, и т.д.
   CREATE TYPE class_requirement AS ENUM ('warrior','mage','monster','any');
   
   
    CREATE TABLE users
    (
        id serial not null unique,
        name varchar(255) not null,
        username varchar(255) not null unique,
        password_hash varchar(255) not null,
        role user_role not null --admin/GM/player
    );

    CREATE TABLE characters
    (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        nickname VARCHAR(255) NOT NULL unique,
        class character_class NOT NULL,
        level INT DEFAULT 1,
        exp INT DEFAULT 0,
        health INT DEFAULT 20,
        strength INT DEFAULT 8,
        agility INT DEFAULT 8,
        charisma INT DEFAULT 8,
        intelligence INT DEFAULT 8,
        created_time TIMESTAMP DEFAULT NOW()
     );

    CREATE TABLE  skills
    (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL unique,
        description TEXT,
        skill_type VARCHAR(255), --passive,active
        effect JSONB,
        required_level int DEFAULT 1,
        required_class class_requirement DEFAULT 'any'
    );

    CREATE TABLE character_skills
    (
        character_id INT REFERENCES characters(id) ON DELETE CASCADE,
        skill_id INT REFERENCES skills(id) ON DELETE CASCADE,
        level INT DEFAULT 1,
        exp INT DEFAULT 0,
        PRIMARY KEY(character_id,skill_id)
    );

    CREATE TABLE items
    (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL unique,
        price INT DEFAULT 1,
        enchantment_level INT DEFAULT 0,
        description TEXT,
        rarity VARCHAR(20) DEFAULT 'common', --COMMON,RARE,EPIC,LEGENDARY,SECRET
        effect JSONB,
        required_level INT DEFAULT 1,
        required_class class_requirement DEFAULT 'any',
        required_strength INT DEFAULT 5,
        required_agility INT DEFAULT 5,
        required_intellegence INT DEFAULT 5,
        required_charisma INT DEFAULT 5
    );

    CREATE TABLE character_inventory
    (
        character_id INT REFERENCES characters(id) ON DELETE CASCADE,
        item_id INT REFERENCES items(id) ON DELETE CASCADE,
        quantity INT DEFAULT 1,
        equipped BOOLEAN DEFAULT FALSE,
        PRIMARY KEY (character_id, item_id)
    );

    CREATE TABLE monsters
    (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        lvl INT DEFAULT 1,
        hp INT DEFAULT  20,
        damage INT DEFAULT 1,
        skills INT REFERENCES skills(id) ON DELETE CASCADE,
        reward_items INT REFERENCES items(id) ON DELETE CASCADE,
        reward_gold INT DEFAULT 1,
        reward_exp INT DEFAULT 10
    );

 
