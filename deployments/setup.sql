CREATE DATABASE IF NOT EXISTS userRepository;

USE userRepository;

CREATE TABLE IF NOT EXISTS userRepository(
				username VARCHAR(50),
				password TEXT NOT NULL,
                firstname TEXT NOT NULL,
                lastname TEXT NOT NULL,
                age INTEGER NOT NULL,
   				gender TEXT NOT NULL,
   				city TEXT NOT NULL,
				country TEXT NULL,
				phone VARCHAR(10) NOT NULL,
				email TEXT NOT NULL,
				githubUsername TEXT NULL,
				PRIMARY KEY (username)
				);


CREATE TABLE IF NOT EXISTS task(
		username VARCHAR(50) NOT NULL,
		name TEXT NOT NULL,
    	description TEXT,
    	start datetime NOT NULL,
    	end   datetime NOT NULL,
    	urlLink   TEXT NULL
		);