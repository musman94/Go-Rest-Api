# createCandidate

Add a new candidate to the database.

**URL** : `/candidate/createCandidate`

**Method** : `POST`

**Data constraints**

Provide all the required information of the Candidate model.

```json
{
	"first_name" : "[string]",
	"last_name"  : "[string]",
	"email"      : "[string]",
	"department" : "[string]",
	"university" : "[string]",
	"experience" : "[bool]",
	"assignee"   : "[string]"
}
```

**Example** All fields must be sent.

```json
{
	"first_name" : "mehmet",
	"last_name"  : "halit",
	"email"      : "mehmethalit@gmail.com",
	"department" : "Development",
	"university" : "Bilkent University",
	"experience" : true,
	"assignee"   : "5e9b2e37a358a83e95401984"
}
```

## Success Response

**Condition** : If everything is OK.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "_id": "5e9b564ab0d66741ab30fed5",
    "first_name": "mehmet",
    "last_name": "halit",
    "email": "mehmethalit@gmail.com",
    "department": "Development",
    "university": "Bilkent University",
    "experience": true,
    "status": "Pending",
    "meeting_count": 0,
    "next_meeting": null,
    "assignee": "5e9b2e37a358a83e95401984",
    "application_date": "18-04-2020"
}
```

## Error Responses

**Condition** : If the email is not valid.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "error": "Invalid email"
}

```

### Or

**Condition** : If the department is not Marketing, Design or Development

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "The department applied to does not exist"
}

```

### Or

**Condition** : If the assignee doesnot exist in the database

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "Assignee not found"
}

```

### Or

**Condition** : If the assignee's department is not the same as the one applied to

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "Assignee should belong to the department the candidate applied to"
}

```

### Or

**Condition** : If there was a problem on the database side

**Code** : `500 BAD REQUEST`

**Content example**

```json
{
    "error": "There was an internal problem while creating the candidate"
}

```
