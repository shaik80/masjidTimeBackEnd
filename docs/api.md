
### Masjid API:

1. **Create a Masjid:**
   - **Endpoint:** `POST /masjids`
   - **Description:** Create a new masjid.
   - **Request Body:**
     ```json
     {
       "name": "Masjid-E-Sahaba",
       "lat": "12.8323672",
       "long": "77.6608719",
       "email": "sh.mudassir98@gmail.com",
       "pass": "hashed_pass",
       "phone_no": "1234567890",
       "country": "India",
       "city_village": "Bangalore",
       "address": "address",
       "status": "active"
     }
     ```
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "Masjid-E-Sahaba",
       "lat": "12.8323672",
       "long": "77.6608719",
       "email": "sh.mudassir98@gmail.com",
       "pass": "hashed_pass",
       "phone_no": "1234567890",
       "country": "India",
       "city_village": "Bangalore",
       "address": "address",
       "status": "active"
     }
     ```

2. **Get Masjid:**
   - **Endpoint:** `GET /masjids/{id}`
   - **Description:** Get details of a specific masjid.
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "Masjid-E-Sahaba",
       "lat": "12.8323672",
       "long": "77.6608719",
       "email": "sh.mudassir98@gmail.com",
       "pass": "hashed_pass",
       "phone_no": "1234567890",
       "country": "India",
       "city_village": "Bangalore",
       "address": "address",
       "status": "active"
     }
     ```

3. **Update Masjid:**
   - **Endpoint:** `PUT /masjids/{id}`
   - **Description:** Update details of a specific masjid.
   - **Request Body:**
     ```json
     {
       "name": "Updated Masjid Name",
       "status": "inactive"
       // other fields to update
     }
     ```
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "Updated Masjid Name",
       // other updated fields
     }
     ```

4. **Delete Masjid:**
   - **Endpoint:** `DELETE /masjids/{id}`
   - **Description:** Delete a specific masjid.
   - **Response:**
     ```json
     {
       "message": "Masjid deleted successfully"
     }
     ```

### Prayer Time API:

5. **Create a Prayer Time:**
   - **Endpoint:** `POST /prayer-times`
   - **Description:** Create a new prayer time.
   - **Request Body:**
     ```json
     {
       "name": "ashar",
       "azan": "1:30 pm",
       "iqama": "1:45 pm"
     }
     ```
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "ashar",
       "azan": "1:30 pm",
       "iqama": "1:45 pm"
     }
     ```

6. **Get Prayer Time:**
   - **Endpoint:** `GET /prayer-times/{id}`
   - **Description:** Get details of a specific prayer time.
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "ashar",
       "azan": "1:30 pm",
       "iqama": "1:45 pm"
     }
     ```

7. **Update Prayer Time:**
   - **Endpoint:** `PUT /prayer-times/{id}`
   - **Description:** Update details of a specific prayer time.
   - **Request Body:**
     ```json
     {
       "name": "Updated Prayer Time Name",
       "azan": "2:00 pm",
       // other fields to update
     }
     ```
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "Updated Prayer Time Name",
       // other updated fields
     }
     ```

8. **Delete Prayer Time:**
   - **Endpoint:** `DELETE /prayer-times/{id}`
   - **Description:** Delete a specific prayer time.
   - **Response:**
     ```json
     {
       "message

": "Prayer time deleted successfully"
     }
     ```

9. **Get Masjid with Array of Prayers:**
   - **Endpoint:** `GET /masjids/{id}/prayers`
   - **Description:** Get details of a specific masjid along with an array of associated prayers.
   - **Response:**
     ```json
     {
       "id": 1,
       "name": "Masjid-E-Sahaba",
       "lat": "12.8323672",
       "long": "77.6608719",
       "email": "sh.mudassir98@gmail.com",
       "phone_no": "1234567890",
       "country": "India",
       "city_village": "Bangalore",
       "address": "address",
       "status": "active",
       "prayers": [
         {
           "id": 1,
           "name": "ashar",
           "azan": "1:30 pm",
           "iqama": "1:45 pm"
         },
         {
           "id": 2,
           "name": "fajar",
           "azan": "5:30 am",
           "iqama": "5:45 am"
         },
         // Add more prayers as needed
       ]
     }
     ```