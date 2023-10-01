import pyodbc

conn = pyodbc.connect(
    "Driver={ODBC Driver 13 for SQL Server};"
    "Server=localhost;"
    "Database=testdb;"
    "UID=sa;"
    "PWD=yourpassword;"
    "Port=1433;"
)

cursor = conn.cursor()

print("Successfully connected to SQL server!")
