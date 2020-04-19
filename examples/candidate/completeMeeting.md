# completeMeeting

Complete a meeting with a candidate

**URL** : `/candidate/completeMeeting/{_id}`

**Method** : `PATCH`

**Data constraints**

Provide the id of the candidate in the url.

**Example** 

```
 /assignee/completeMeeting/5e9b5ec12608343d5cbf8c77"`

```

## Success Response

**Condition** : If everything is OK and the candidate exists in the database.

**Code** : `200 SUCCESS`

**Content example**

```json
{
    "Meeting completed with candidate": "5e9b5ec12608343d5cbf8c77"
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
