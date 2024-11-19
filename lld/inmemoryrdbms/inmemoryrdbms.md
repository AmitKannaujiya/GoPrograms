**Problem :** Designing an in-memory database system requires careful consideration of data structures and operations to support core database features like constraints, type checking, and CRUD operations. Below is the Low-Level Design (LLD) for such a system.

**Features**
**Create Database:** Define a new in-memory database.
**Drop Database:** Delete an existing database.
**Create Table:** Define a new table with schema, types, and constraints.
**Drop Table**: Delete a table and its data.
**Truncate Table:** Clear all rows in a table.
**Insert:** Add a row to a table with type and constraint checks.
**Delete:** Remove rows matching specific conditions.
**Size and Type Check:** Ensure column values match defined types and constraints.
**Query Operations:** Simple operations like fetching rows (basic SELECT).

System Design
1. Classes/Components
**DatabaseManager**:
Handles database creation, deletion, and switching between databases.
**Table:**
Represents a table schema and its rows.
**Row:**
Represents a single record in a table.
**Schema:**
Defines column names, types, and constraints for a table.
**DataTypes and Constraints:**
Type system (e.g., INT, STRING) and constraints (e.g., UNIQUE, NOT NULL).
**QueryProcessor:**
Processes commands like INSERT, DELETE, SELECT, etc.
