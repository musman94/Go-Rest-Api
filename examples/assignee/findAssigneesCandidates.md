# Get all of the assignee's candidates

Send a message from one user to another.

**URL** : `/assignee/findAssigneesCandidates/{_id}`

**Method** : `GET`

**Data constraints**

Provide the name of the assignee in the url.

**Example** 

```
/assignee/findAssigneesCandidates/5e9c4a78323cc69c4f34da78`
```

## Success Response

**Condition** : If everything is OK and the assignee exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Candidates": [
        {
            "_id": "5b75881051d9590001def62a",
            "first_name": "Kadir",
            "last_name": "kurt",
            "email": "kadirkurt@student.com",
            "department": "Design",
            "university": "METU",
            "experience": false,
            "status": "In Progress",
            "meeting_count": 3,
            "next_meeting": "2020-05-01T12:40:00Z",
            "assignee": "5c191acea7948900011168d4",
            "application_date": ""
        },
        {
            "_id": "5b75865c51d9590001def626",
            "first_name": "Ayşe",
            "last_name": "Kaya",
            "email": "ayşekaya@yaho.com",
            "department": "Design",
            "university": "Ankara",
            "experience": false,
            "status": "Accepted",
            "meeting_count": 4,
            "next_meeting": null,
            "assignee": "5c191acea7948900011168d4",
            "application_date": ""
        }
    ]
}

```

## Error Response

**Condition** : If the assignee doesn't exist in the database.

**Code** : `404 NOT FOUND`

**Content example** :

```json
{
    "error": "Invalid assignee ID, assignee couldn't be found""
}
```
