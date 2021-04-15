Simple Application Backend

Host : localhost:8000
Endpoint : /api/log
Request Method : POST
Output : server.log dan http.StatusOK
Header : X-RANDOM

Request data (JSON)
{
   "counter": 1
}

Input program yaitu header X-RANDOM dan data JSON
Output program yaitu memberikan status 201 kepada client dan menghasilkan file server.log