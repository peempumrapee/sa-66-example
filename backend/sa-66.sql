BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "members" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"username"	text,
	"password"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "creators" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"username"	text,
	"password"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "sound_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "sounds" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"title"	text,
	"path"	text,
	"sound_type_id"	integer,
	"creator_id"	integer,
	CONSTRAINT "fk_sound_types_sounds" FOREIGN KEY("sound_type_id") REFERENCES "sound_types"("id"),
	CONSTRAINT "fk_creators_sounds" FOREIGN KEY("creator_id") REFERENCES "creators"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "histories" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"played_at"	datetime,
	"sound_id"	integer,
	"member_id"	integer,
	CONSTRAINT "fk_sounds_histories" FOREIGN KEY("sound_id") REFERENCES "sounds"("id"),
	CONSTRAINT "fk_members_histories" FOREIGN KEY("member_id") REFERENCES "members"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "ratings" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"score"	integer,
	"sound_id"	integer,
	"member_id"	integer,
	CONSTRAINT "fk_sounds_ratings" FOREIGN KEY("sound_id") REFERENCES "sounds"("id"),
	CONSTRAINT "fk_members_ratings" FOREIGN KEY("member_id") REFERENCES "members"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "playlists" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"title"	text,
	"member_id"	integer,
	CONSTRAINT "fk_members_playlists" FOREIGN KEY("member_id") REFERENCES "members"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "sound_playlists" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"playlist_id"	integer,
	"sound_id"	integer,
	CONSTRAINT "fk_playlists_sound_playlists" FOREIGN KEY("playlist_id") REFERENCES "playlists"("id"),
	CONSTRAINT "fk_sounds_sound_playlists" FOREIGN KEY("sound_id") REFERENCES "sounds"("id"),
	PRIMARY KEY("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_members_username" ON "members" (
	"username"
);
CREATE INDEX IF NOT EXISTS "idx_members_deleted_at" ON "members" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_members_email" ON "members" (
	"email"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_creators_email" ON "creators" (
	"email"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_creators_username" ON "creators" (
	"username"
);
CREATE INDEX IF NOT EXISTS "idx_creators_deleted_at" ON "creators" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_sound_types_deleted_at" ON "sound_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_sounds_deleted_at" ON "sounds" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_histories_deleted_at" ON "histories" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_ratings_deleted_at" ON "ratings" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_playlists_deleted_at" ON "playlists" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_sound_playlists_deleted_at" ON "sound_playlists" (
	"deleted_at"
);
COMMIT;