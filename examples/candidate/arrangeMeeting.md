# arrangeMeeting

Arrange a meeting with a candidate

**URL** : `/candidate/arrangeMeeting/{_id}`

**Method** : `PATCH`

**Data constraints**

Provide the id of the candidate in the url.

Provide the meeting time in the body. 

```json
{
	"next_meeting" : "[*time.Time]"
}
```

**Example** 

```
 /assignee/arrangeMeeting/5e9b5ec12608343d5cbf8c77"`

 ```json
{
	"next_meeting" : "0004-01-01T00:00:00Z"
}
```

```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Meeting arranged with candidate": "5e9b5ec12608343d5cbf8c77"
}
```

## Error Response

**Condition** : If the candidate doesn't exist in the database.

**Code** : `404 NOT FOUND`

**Content example** :

```json
{
    "error": "Invalid Candidate ID"
}
```
