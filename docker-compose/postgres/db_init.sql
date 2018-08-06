--
-- User creation
--

CREATE USER testuser
    WITH PASSWORD 'md599e5ea7a6f7c3269995cba3927fd0093';

--
-- Database creation
--

CREATE DATABASE testdb
    WITH OWNER testuser;

--
-- Access rights
--

REVOKE ALL ON DATABASE testdb FROM PUBLIC;
GRANT ALL ON DATABASE testdb TO testuser;