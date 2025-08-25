package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Action was created
	Person EmbeddedResource `json:"person,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Action was last updated
	Multitouch_group_id int `json:"multitouch_group_id,omitempty"` // ID of the multitouch group
	Status string `json:"status,omitempty"` // The current state of the person on the cadence. Possible values are: in_progress: this action has not been completed pending_activity: this action has been acted upon, but the action has not been completed. (i.e. the email is scheduled to send, but has not been delivered yet)
	User EmbeddedResource `json:"user,omitempty"`
	Action_details EmbeddedResource `json:"action_details,omitempty"`
	Id int `json:"id,omitempty"` // ID of Action
	Step EmbeddedResource `json:"step,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of this action. Valid types are: email, phone, other. New types may be added in the future.
	Due bool `json:"due,omitempty"` // Whether this step is due
	Due_on string `json:"due_on,omitempty"` // When action is due
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Activity was last updated
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Id int `json:"id,omitempty"` // Import ID
	Imported_people_count int `json:"imported_people_count,omitempty"` // Count of People that have ever been on this Import
	Name string `json:"name,omitempty"` // Name of Import
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the import was last updated, ignoring relationship changes
	Created_at string `json:"created_at,omitempty"` // Datetime of when the import was created
	Current_people_count int `json:"current_people_count,omitempty"` // Count of People that have not been deleted
}

// Step represents the Step schema from the OpenAPI specification
type Step struct {
	Details EmbeddedResource `json:"details,omitempty"`
	Disabled bool `json:"disabled,omitempty"` // Whether this step is currently active
	Name string `json:"name,omitempty"` // Name of the step
	Step_number int `json:"step_number,omitempty"` // The number of the step for this day
	TypeField string `json:"type,omitempty"` // The type of the action scheduled by this step. Valid types are: email, phone, integration, other. New types may be added in the future.
	Id int `json:"id,omitempty"` // ID of Step
	Multitouch_enabled bool `json:"multitouch_enabled,omitempty"` // Whether this step is a multitouch cadence step
	Day int `json:"day,omitempty"` // Day this step is associated with up
	Display_name string `json:"display_name,omitempty"` // Display name of the step
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Step was last updated
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Step was created
}

// EventMeetingSetting represents the EventMeetingSetting schema from the OpenAPI specification
type EventMeetingSetting struct {
	Email_address string `json:"email_address,omitempty"` // Calendar owner's email address
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Sentiments_required bool `json:"sentiments_required,omitempty"` // Whether team members are required to log sentiments
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the team was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the team was created
	Local_dial_enabled bool `json:"local_dial_enabled,omitempty"` // Whether this team has local dial enabled
	License_limit int `json:"license_limit,omitempty"` // Count of seats that this team has licensed
	Click_tracking_default bool `json:"click_tracking_default,omitempty"` // The team default for click tracking when composing emails
	Team_visibility_default string `json:"team_visibility_default,omitempty"` // The default visibility of resources on the team, in the UI only. The API does not utilize this default. Possible values are: public, private.
	Id int `json:"id,omitempty"` // Team ID
	Name string `json:"name,omitempty"` // Team name
	Plan string `json:"plan,omitempty"` // Plan type of the team, Possible values are: group, professional, enterprise
	Record_by_default bool `json:"record_by_default,omitempty"` // Whether calls will record by default
	Allow_automated_email_steps bool `json:"allow_automated_email_steps,omitempty"` // Whether team members are allowed to have automated email steps
	Plan_features map[string]interface{} `json:"plan_features,omitempty"` // Add on features for this team
	Deactivated bool `json:"deactivated,omitempty"` // Indicates if the team has been deactivated
	Email_daily_limit int `json:"email_daily_limit,omitempty"` // Daily email limit for each member on the team
	Private_fields map[string]interface{} `json:"_private_fields,omitempty"` // For internal use only. This field does not comply with our backwards compatability policies.
	Dispositions_required bool `json:"dispositions_required,omitempty"` // Whether team members are required to mark disposition at the end of calls
	Custom_tracking_domain string `json:"custom_tracking_domain,omitempty"` // The domain click and open tracking will be proxied through
	Group_privacy_setting string `json:"group_privacy_setting,omitempty"` // Visibility setting for resources across the team. Possible values are: group_public, all_public. When the value is group_public, certain resources will only be visible to members of the same group. When the value is all_public, all resources are visible to all users on this team.
	Call_recording_disabled bool `json:"call_recording_disabled,omitempty"` // Whether all call recording is disabled
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Id int `json:"id,omitempty"` // ID of Custom Field
	Name string `json:"name,omitempty"` // Name of the Custom Field
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Custom Field was last updated
	Value_type string `json:"value_type,omitempty"` // Value Type of the Custom Field. Value must be one of: text, date.
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Custom Field was created
	Field_type string `json:"field_type,omitempty"` // Type of the Custom Field. Value must be one of: person, company, opportunity.
}

// ConversationsCall represents the ConversationsCall schema from the OpenAPI specification
type ConversationsCall struct {
	Call_created_at string `json:"call_created_at,omitempty"` // Timestamp for when the call started. If not provided, will default to the time the request was received
	Direction string `json:"direction,omitempty"` // Call direction
	Duration float64 `json:"duration,omitempty"` // Duration of call in seconds
	From string `json:"from,omitempty"` // Phone number that call was made from
	Recording map[string]interface{} `json:"recording,omitempty"` // Object containing recording info including the audio file (.mp3, .wav, .ogg, .m4a)
	To string `json:"to,omitempty"` // Phone number that was called
	User_guid string `json:"user_guid,omitempty"` // Guid of the Salesloft User to assign the call to. If not provided, will default to the user within the authentication token
}

// CallSentiment represents the CallSentiment schema from the OpenAPI specification
type CallSentiment struct {
	Name string `json:"name,omitempty"` // An available call sentiment text
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the call sentiment was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the call sentiment was created
	Id int `json:"id,omitempty"` // ID of CallSentiment
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Do_not_contact bool `json:"do_not_contact,omitempty"` // Whether this company has opted out of communications. Do not contact someone at this company when this is set to true
	Counts EmbeddedAccountCounts `json:"counts,omitempty"`
	Id int `json:"id,omitempty"` // ID of Account
	Country string `json:"country,omitempty"` // Country
	Description string `json:"description,omitempty"` // Description
	Locale string `json:"locale,omitempty"` // Time locale
	Industry string `json:"industry,omitempty"` // Industry
	Phone string `json:"phone,omitempty"` // Phone number without formatting
	State string `json:"state,omitempty"` // State
	Owner_crm_id string `json:"owner_crm_id,omitempty"` // Mapped owner field from the CRM
	Revenue_range string `json:"revenue_range,omitempty"` // Estimated revenue range
	Twitter_handle string `json:"twitter_handle,omitempty"` // Twitter handle, with @
	Size string `json:"size,omitempty"` // Estimated number of people in employment
	User_relationships map[string]interface{} `json:"user_relationships,omitempty"` // Filters by accounts matching all given user relationship fields, _is_null or _unmapped can be passed to filter accounts with null or unmapped user relationship values
	Company_stage EmbeddedResource `json:"company_stage,omitempty"`
	Last_contacted_person EmbeddedResource `json:"last_contacted_person,omitempty"`
	Domain string `json:"domain,omitempty"` // Website domain, not a fully qualified URI
	Crm_id string `json:"crm_id,omitempty"` // CRM ID
	Owner EmbeddedResource `json:"owner,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"` // Custom fields are defined by the user's team. Only fields with values are presented in the API.
	Street string `json:"street,omitempty"` // Street name and number
	Account_tier EmbeddedResource `json:"account_tier,omitempty"`
	Website string `json:"website,omitempty"` // Website
	Crm_object_type string `json:"crm_object_type,omitempty"` // CRM object type
	Name string `json:"name,omitempty"` // Account Full Name
	Postal_code string `json:"postal_code,omitempty"` // Postal code
	Archived_at string `json:"archived_at,omitempty"` // Datetime of when the Account was archived, if archived
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Account was created
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Account was last updated
	Last_contacted_by EmbeddedResource `json:"last_contacted_by,omitempty"`
	Linkedin_url string `json:"linkedin_url,omitempty"` // Full LinkedIn url
	Conversational_name string `json:"conversational_name,omitempty"` // Conversational name of the Account
	Founded string `json:"founded,omitempty"` // Date or year of founding
	Last_contacted_type string `json:"last_contacted_type,omitempty"` // The type of the last touch to this Account. Can be call, email, other
	Tags []string `json:"tags,omitempty"` // All tags applied to this Account
	Company_type string `json:"company_type,omitempty"` // Type of the Account's company
	Last_contacted_at string `json:"last_contacted_at,omitempty"` // Datetime this Account was last contacted
	City string `json:"city,omitempty"` // City
	Creator EmbeddedResource `json:"creator,omitempty"`
	Crm_url string `json:"crm_url,omitempty"` // CRM url
}

// MeetingSetting represents the MeetingSetting schema from the OpenAPI specification
type MeetingSetting struct {
	User_slug string `json:"user_slug,omitempty"` // User slug generated with a full name of the user
	User EmbeddedResource `json:"user,omitempty"`
	Primary_calendar_id string `json:"primary_calendar_id,omitempty"` // ID of the primary calendar
	Availability_limit_enabled bool `json:"availability_limit_enabled,omitempty"` // If Availability Limits have been turned on
	Schedule_buffer_enabled bool `json:"schedule_buffer_enabled,omitempty"` // Determines if meetings are scheduled with a 15 minute buffer between them
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the MeetingSetting was last updated
	Reschedule_meetings_enabled bool `json:"reschedule_meetings_enabled,omitempty"` // Determines if a user enabled reschedule meetings feature
	Enable_calendar_sync bool `json:"enable_calendar_sync,omitempty"` // Determines if a user enabled Calendar Sync feature
	Allow_booking_overtime bool `json:"allow_booking_overtime,omitempty"` // Allow team members to insert available time outside your working hours.
	Primary_calendar_connection_failed bool `json:"primary_calendar_connection_failed,omitempty"` // Gets true when any issue with fetching calendar occurs
	Enable_dynamic_location bool `json:"enable_dynamic_location,omitempty"` // Determines if location will be filled via third-party service (Zoom, GoToMeeting, etc.)
	Buffer_time_duration int `json:"buffer_time_duration,omitempty"` // Default buffer duration in minutes set by a user
	Share_event_detail bool `json:"share_event_detail,omitempty"` // Allow team members to see the details of events on your calendar.
	Description string `json:"description,omitempty"` // Default description of the meeting
	Primary_calendar_name string `json:"primary_calendar_name,omitempty"` // Display name of the primary calendar
	Allow_booking_on_behalf bool `json:"allow_booking_on_behalf,omitempty"` // Allow other team members to schedule on you behalf.
	Calendar_type string `json:"calendar_type,omitempty"` // Calendar type
	Id int `json:"id,omitempty"` // ID of the MeetingSetting
	Created_at string `json:"created_at,omitempty"` // Datetime of when the MeetingSetting was created
	Schedule_delay int `json:"schedule_delay,omitempty"` // The number of hours in advance a user requires someone to a book a meeting with them
	Times_available map[string]interface{} `json:"times_available,omitempty"` // Times available set by a user that can be used to book meetings
	Title string `json:"title,omitempty"` // Default title of the meeting
	Time_zone string `json:"time_zone,omitempty"` // Time zone for current calendar
	Allow_event_overlap bool `json:"allow_event_overlap,omitempty"` // Allow team members to double book events on your calendar.
	Availability_limit int `json:"availability_limit,omitempty"` // The number of days out the user allows a prospect to schedule a meeting
	User_details map[string]interface{} `json:"user_details,omitempty"` // User details
	Location string `json:"location,omitempty"` // Default location of the meeting
	Active_meeting_url MeetingUrl `json:"active_meeting_url,omitempty"`
	Email_address string `json:"email_address,omitempty"` // Calendar owner's email address
	Default_meeting_length int `json:"default_meeting_length,omitempty"` // Default meeting length in minutes set by the user
}

// CrmActivityField represents the CrmActivityField schema from the OpenAPI specification
type CrmActivityField struct {
	Value string `json:"value,omitempty"` // A value to always be written. This value does not need to be sent to other endpoints' crm params, but must be the exact value if sent. Email source fields will always have a value present.
	Field string `json:"field,omitempty"` // The CRM field name
	Picklist_values map[string]interface{} `json:"picklist_values,omitempty"` // Valid picklist values, if present for this field. The format is {label => value}. If present, only values in the picklist structure can be used as a crm param.
	Source string `json:"source,omitempty"` // SalesLoft object that this field is mapped for. Valid sources are: email, phone
	Field_type string `json:"field_type,omitempty"` // The type of this field in your CRM. Certain field types can only accept structured input.
	Id int `json:"id,omitempty"` // ID of CrmActivityField
	Salesforce_object_type string `json:"salesforce_object_type,omitempty"` // The Salesforce object type that this field maps to. Valid object types are: Task. More object types may be added in the future.
	Title string `json:"title,omitempty"` // A human friendly title for this field
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the CrmActivityField was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the CrmActivityField was created
	Crm_object_type string `json:"crm_object_type,omitempty"` // The CRM object type that this field maps to. Valid object types are CRM dependent: Task, Phonecall, Email.
}

// ActivityHistory represents the ActivityHistory schema from the OpenAPI specification
type ActivityHistory struct {
	Occurred_at string `json:"occurred_at,omitempty"` // When this activity occurred
	Resource_type int `json:"resource_type,omitempty"` // Type of the resource this activity is for. One of: account, person
	TypeField string `json:"type,omitempty"` // The type of activity
	Static_data map[string]interface{} `json:"static_data,omitempty"` // The static data for this activity
	User_guid string `json:"user_guid,omitempty"` // UUID of the user this activity is for
	Dynamic_data map[string]interface{} `json:"dynamic_data,omitempty"` // Attributes from associated records. This is specific to the type of activity and may change over time. Not returned for create requests
	Failed_dynamic_resources map[string]interface{} `json:"failed_dynamic_resources,omitempty"` // A list of remote resource names that failed to load. This is specific to the type of activity and may change over time. Not returned for create requests
	Pinned_at string `json:"pinned_at,omitempty"` // When this record was pinned
	Updated_at string `json:"updated_at,omitempty"` // When this record was updated
	Created_at string `json:"created_at,omitempty"` // When this record was created
	Id int `json:"id,omitempty"` // ID of this activity
	Resource_id int `json:"resource_id,omitempty"` // ID of the resource this activity is for. It will be a string for the following resource types: crm_opportunity
}

// PersonStage represents the PersonStage schema from the OpenAPI specification
type PersonStage struct {
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Person Stage was created
	Id int `json:"id,omitempty"` // ID of Person Stage
	Name string `json:"name,omitempty"` // Name of Person Stage
	Order int `json:"order,omitempty"` // Sortable value of Person Stage order
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Person Stage was last updated
}

// EmailCounts represents the EmailCounts schema from the OpenAPI specification
type EmailCounts struct {
	Attachments int `json:"attachments,omitempty"` // The number of attachments on the email
	Clicks int `json:"clicks,omitempty"` // The number of times links in the email were clicked
	Replies int `json:"replies,omitempty"` // The number of replies the email received
	Unique_devices int `json:"unique_devices,omitempty"` // The number of unique devices that opened the email
	Unique_locations int `json:"unique_locations,omitempty"` // The number of unique locations that opened the email
	Views int `json:"views,omitempty"` // The number of times the email was opened
}

// Person represents the Person schema from the OpenAPI specification
type Person struct {
	Linkedin_url string `json:"linkedin_url,omitempty"` // Linkedin URL
	Account EmbeddedResource `json:"account,omitempty"`
	Personal_website string `json:"personal_website,omitempty"` // The website of this person
	Display_name string `json:"display_name,omitempty"` // Either the full name or the email address. Use this when showing a person's name
	Created_at string `json:"created_at,omitempty"` // Datetime of when the person was created
	Crm_url string `json:"crm_url,omitempty"` // CRM url
	Untouched bool `json:"untouched,omitempty"` // The person's untouched status
	Counts PersonCounts `json:"counts,omitempty"`
	Last_replied_at string `json:"last_replied_at,omitempty"` // Last datetime this person replied to an email
	Starred bool `json:"starred,omitempty"` // Whether this person is starred by the current user
	Job_seniority string `json:"job_seniority,omitempty"` // The Job Seniority of a Person, must be one of director, executive, individual_contributor, manager, vice_president, unknown
	Work_city string `json:"work_city,omitempty"` // Work location - city
	Email_address string `json:"email_address,omitempty"` // Email address
	Locale_utc_offset int `json:"locale_utc_offset,omitempty"` // The locale's timezone offset from UTC in minutes
	Most_recent_cadence EmbeddedResource `json:"most_recent_cadence,omitempty"`
	Secondary_email_address string `json:"secondary_email_address,omitempty"` // Alternate email address
	Person_stage EmbeddedResource `json:"person_stage,omitempty"`
	Title string `json:"title,omitempty"` // Job title
	Id int `json:"id,omitempty"` // Person ID
	Mobile_phone string `json:"mobile_phone,omitempty"` // Mobile phone without formatting
	Country string `json:"country,omitempty"` // Country
	First_name string `json:"first_name,omitempty"` // First name
	Owner_crm_id string `json:"owner_crm_id,omitempty"` // Mapped owner field from your CRM
	Cadences []EmbeddedResource `json:"cadences,omitempty"` // The list of active cadences person is added to
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"` // Custom fields are defined by the user's team. Only fields with values are presented in the API.
	Locale string `json:"locale,omitempty"` // Time locale of the person
	Person_company_name string `json:"person_company_name,omitempty"` // Company name. This property is specific to this person, unrelated to the company object. Updating the company object associated with this person is recommended
	City string `json:"city,omitempty"` // City
	Phone string `json:"phone,omitempty"` // Phone without formatting
	Work_country string `json:"work_country,omitempty"` // Work location - country
	Work_state string `json:"work_state,omitempty"` // Work location - state
	Last_contacted_type string `json:"last_contacted_type,omitempty"` // The type of the last touch to this person. Can be call, email, other
	Crm_object_type string `json:"crm_object_type,omitempty"` // CRM object type
	Home_phone string `json:"home_phone,omitempty"` // Home phone without formatting
	State string `json:"state,omitempty"` // State
	Crm_id string `json:"crm_id,omitempty"` // CRM ID
	Last_contacted_at string `json:"last_contacted_at,omitempty"` // Last datetime this person was contacted
	Last_contacted_by EmbeddedResource `json:"last_contacted_by,omitempty"`
	Success_count int `json:"success_count,omitempty"` // The person's success count. 1 if person has any active successes, 0 otherwise.
	Contact_restrictions []string `json:"contact_restrictions,omitempty"` // Specific methods of communication to prevent for this person. This will prevent individual execution of these communication types as well as automatically skip cadence steps of this communication type for this person in SalesLoft. Values currently accepted: call, email, message
	Do_not_contact bool `json:"do_not_contact,omitempty"` // Whether or not this person has opted out of all communication. Setting this value to true prevents this person from being called, emailed, or added to a cadence in SalesLoft. If this person is currently in a cadence, they will be removed.
	Last_completed_step EmbeddedResource `json:"last_completed_step,omitempty"`
	Bouncing bool `json:"bouncing,omitempty"` // Whether this person's current email address has bounced
	Person_company_industry string `json:"person_company_industry,omitempty"` // Company industry. This property is specific to this person, unrelated to the company object. Updating the company object associated with this person is recommended
	Owner EmbeddedResource `json:"owner,omitempty"`
	Tags []string `json:"tags,omitempty"` // All tags applied to this person
	Person_company_website string `json:"person_company_website,omitempty"` // Company website. This property is specific to this person, unrelated to the company object. Updating the company object associated with this person is recommended
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the person was last updated
	ImportField EmbeddedResource `json:"import,omitempty"`
	Last_completed_step_cadence EmbeddedResource `json:"last_completed_step_cadence,omitempty"`
	Eu_resident bool `json:"eu_resident,omitempty"` // Whether this person is marked as a European Union Resident or not
	Personal_email_address string `json:"personal_email_address,omitempty"` // Personal email address
	Phone_extension string `json:"phone_extension,omitempty"` // Phone extension without formatting
	Twitter_handle string `json:"twitter_handle,omitempty"` // The twitter handle of this person
	Last_name string `json:"last_name,omitempty"` // Last name
	Full_email_address string `json:"full_email_address,omitempty"` // Full email address with name
}

// CalendarEvent represents the CalendarEvent schema from the OpenAPI specification
type CalendarEvent struct {
	I_cal_uid string `json:"i_cal_uid,omitempty"` // Calendar event unique identifier (iCalUID)
	Status string `json:"status,omitempty"` // The status of the calendar event. It can be empty for non-google events.
	All_day bool `json:"all_day,omitempty"` // Whether the calendar event is an all-day event.
	Tenant_id int `json:"tenant_id,omitempty"` // Tenant ID of the user calendar
	Canceled_at string `json:"canceled_at,omitempty"` // The canceled date of the calendar event.
	Title string `json:"title,omitempty"` // Title of the calendar event
	End_time string `json:"end_time,omitempty"` // The (exclusive) end time of the calendar event.
	Location string `json:"location,omitempty"` // Location of the calendar event
	Start_time string `json:"start_time,omitempty"` // The (inclusive) start time of the calendar event.
	Created_at string `json:"created_at,omitempty"` // Creation time of the calendar event.
	Extended_properties map[string]interface{} `json:"extended_properties,omitempty"` // Extended properties of the calendar event.
	Conference_data map[string]interface{} `json:"conference_data,omitempty"` // The conference-related information, such as details of a Google Meet conference.
	Attendees []interface{} `json:"attendees,omitempty"` // The attendees of the calendar event.
	Calendar_id string `json:"calendar_id,omitempty"` // Calendar ID of the user calendar.
	Html_link string `json:"html_link,omitempty"` // An absolute link to this calendar event in the Google Calendar Web UI.
	Updated_at string `json:"updated_at,omitempty"` // Last modification time of the calendar event.
	Body_html string `json:"body_html,omitempty"` // Raw body content from Microsoft calendar events
	Organizer string `json:"organizer,omitempty"` // The organizer email of the calendar event.
	Description string `json:"description,omitempty"` // Description of the calendar event
	User_guid string `json:"user_guid,omitempty"` // User GUID of the user calendar.
	Id string `json:"id,omitempty"` // The calendar event original ID from calendar provider
	Provider string `json:"provider,omitempty"` // The provider of the calendar event.
	Recurring bool `json:"recurring,omitempty"` // Whether the calendar event is a recurring event.
	Busy bool `json:"busy,omitempty"` // Busy/free status of the calendar event
	Creator string `json:"creator,omitempty"` // The creator email of the calendar event.
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Remind_at string `json:"remind_at,omitempty"` // Datetime of when the user will be reminded of the task, ISO-8601 datetime format required
	User EmbeddedResource `json:"user,omitempty"`
	Current_state string `json:"current_state,omitempty"` // The state of the task. Valid states are: scheduled, completed
	Task_type string `json:"task_type,omitempty"` // The type of the task. Valid types are: call, email, general
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Task was created
	Due_date string `json:"due_date,omitempty"` // Date of when the Task is due, ISO-8601 date format required
	Completed_at string `json:"completed_at,omitempty"` // Datetime of when the task was completed, ISO-8601 datetime format required
	Completed_by EmbeddedResource `json:"completed_by,omitempty"`
	Person EmbeddedResource `json:"person,omitempty"`
	Due_at string `json:"due_at,omitempty"` // Datetime of when the Task is due, can be null. ISO-8601 datetime format required
	Created_by_user EmbeddedResource `json:"created_by_user,omitempty"`
	Description string `json:"description,omitempty"` // A description of the task recorded for person at completion time
	Subject string `json:"subject,omitempty"` // Subject line of the task
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Task was last updated
	Id int `json:"id,omitempty"` // ID of Task
}

// AccountStage represents the AccountStage schema from the OpenAPI specification
type AccountStage struct {
	Name string `json:"name,omitempty"` // Name of Account Stage
	Order int `json:"order,omitempty"` // Order of Account Stage
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Account Stage was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Account Stage was created
	Id int `json:"id,omitempty"` // ID of Account Stage
}

// PersonCounts represents the PersonCounts schema from the OpenAPI specification
type PersonCounts struct {
	Emails_sent int `json:"emails_sent,omitempty"` // The number of emails sent to this person
	Emails_viewed int `json:"emails_viewed,omitempty"` // The number of unique emails viewed by this person
	Calls int `json:"calls,omitempty"` // The number of calls logged to this person
	Emails_bounced int `json:"emails_bounced,omitempty"` // The number of unique emails sent to this person that bounced
	Emails_clicked int `json:"emails_clicked,omitempty"` // The number of unique emails clicked by this person
	Emails_replied_to int `json:"emails_replied_to,omitempty"` // The number of unique emails replied to by this person
}

// LiveFeedItem represents the LiveFeedItem schema from the OpenAPI specification
type LiveFeedItem struct {
	Event_occurred_at string `json:"event_occurred_at,omitempty"` // When this event occurred
	User_guid string `json:"user_guid,omitempty"` // UUID of the user this item is for
	Event_type string `json:"event_type,omitempty"` // The type of event
	Id int `json:"id,omitempty"` // ID of this item
	Message string `json:"message,omitempty"` // A plaintext message for this event
	Alert_metadata map[string]interface{} `json:"alert_metadata,omitempty"` // Information about whether this event should trigger an alert
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The metadata created for this event
	Rollup_key string `json:"rollup_key,omitempty"` // The key that should be used to rollup events client side. null or empty values should not be rolled up
	Title string `json:"title,omitempty"` // A plaintext title for this event
	Path string `json:"path,omitempty"` // The path to the application that should be followed
}

// JobData represents the JobData schema from the OpenAPI specification
type JobData struct {
	Bulk_job map[string]interface{} `json:"bulk_job,omitempty"` // Associated bulk job
	Record map[string]interface{} `json:"record,omitempty"` // The data that was used to process the operation
	Resource map[string]interface{} `json:"resource,omitempty"` // The object containing the resulting resource from performing the bulk action on this record
	ErrorField string `json:"error,omitempty"` // Error associated with this record
	Created_at string `json:"created_at,omitempty"` // When this job data record was created
	Finished_at string `json:"finished_at,omitempty"` // When this job data record finished processing
	Id int `json:"id,omitempty"` // ID of this Job Data
	Started_at string `json:"started_at,omitempty"` // When this job data record started processing
	Status string `json:"status,omitempty"` // Status of this job data. Must be one of: pending, success, error, retrying
}

// CallerId represents the CallerId schema from the OpenAPI specification
type CallerId struct {
	Person EmbeddedResource `json:"person,omitempty"`
	Title string `json:"title,omitempty"` // The title of the person calling
	Account_name string `json:"account_name,omitempty"` // The account of the person calling
	Display_name string `json:"display_name,omitempty"` // The name of the person calling
}

// SavedListView represents the SavedListView schema from the OpenAPI specification
type SavedListView struct {
	Id int `json:"id,omitempty"` // ID of Ssaved list view
	Is_default bool `json:"is_default,omitempty"` // Whether the saved list view is the default view
	Name string `json:"name,omitempty"` // Name of saved list view
	View string `json:"view,omitempty"` // Type of saved list view
	View_params map[string]interface{} `json:"view_params,omitempty"` // List of set filters in saved list view
}

// CallInstruction represents the CallInstruction schema from the OpenAPI specification
type CallInstruction struct {
	Id int `json:"id,omitempty"` // ID of call instructions
	Instructions string `json:"instructions,omitempty"` // The instructions
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the call instructions were last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the call instructions were created
}

// CustomRole represents the CustomRole schema from the OpenAPI specification
type CustomRole struct {
	Id string `json:"id,omitempty"` // ID of the custom role
	Name string `json:"name,omitempty"` // Name of the custom role
}

// CadenceImport represents the CadenceImport schema from the OpenAPI specification
type CadenceImport struct {
	Cadence EmbeddedResource `json:"cadence,omitempty"`
}

// Note represents the Note schema from the OpenAPI specification
type Note struct {
	Associated_with EmbeddedResource `json:"associated_with,omitempty"`
	Call EmbeddedResource `json:"call,omitempty"`
	Content string `json:"content,omitempty"` // The content of the note
	Created_at string `json:"created_at,omitempty"` // Datetime of when the note was created
	Id int `json:"id,omitempty"` // Note ID
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the note was last updated
	User EmbeddedResource `json:"user,omitempty"`
	Associated_type string `json:"associated_type,omitempty"` // Type of associated resource ('person' or 'account')
}

// AccountUpsert represents the AccountUpsert schema from the OpenAPI specification
type AccountUpsert struct {
	Account Account `json:"account,omitempty"`
	Upsert_type string `json:"upsert_type,omitempty"` // The type of upsert. One of: create, update
}

// Cadence represents the Cadence schema from the OpenAPI specification
type Cadence struct {
	Added_stage EmbeddedResource `json:"added_stage,omitempty"`
	Remove_bounces_enabled bool `json:"remove_bounces_enabled,omitempty"` // Whether this cadence is configured to automatically remove people who have bounced
	Groups []EmbeddedResource `json:"groups,omitempty"` // Groups to which this cadence is assigned, if any
	Cadence_function string `json:"cadence_function,omitempty"` // The use case of the cadence. Possible values are: outbound: Denotes an outbound cadence, typically for sales purposes inbound: Denotes an inbound sales cadence event: Denotes a cadence used for an upcoming event other: Denotes a cadence outside of the standard process
	Id int `json:"id,omitempty"` // ID of cadence
	Cadence_priority EmbeddedResource `json:"cadence_priority,omitempty"`
	Opt_out_link_included bool `json:"opt_out_link_included,omitempty"` // Whether this cadence is configured to include an opt-out link by default
	Counts CadenceCounts `json:"counts,omitempty"`
	Team_cadence bool `json:"team_cadence,omitempty"` // Whether this cadence is a team cadence. A team cadence is created by an admin and can be run by all users
	Draft bool `json:"draft,omitempty"` // Whether this cadence is in draft mode
	Finished_stage EmbeddedResource `json:"finished_stage,omitempty"`
	Owner EmbeddedResource `json:"owner,omitempty"`
	Tags []string `json:"tags,omitempty"` // All tags applied to this cadence
	External_identifier string `json:"external_identifier,omitempty"` // Cadence External ID
	Name string `json:"name,omitempty"` // Cadence name
	Replied_stage EmbeddedResource `json:"replied_stage,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the cadence was last updated
	Archived_at string `json:"archived_at,omitempty"` // Datetime of when the cadence was archived, if archived
	Remove_replies_enabled bool `json:"remove_replies_enabled,omitempty"` // Whether this cadence is configured to automatically remove people who have replied
	Bounced_stage EmbeddedResource `json:"bounced_stage,omitempty"`
	Cadence_framework_id int `json:"cadence_framework_id,omitempty"` // ID of the cadence framework used to create steps for the cadence
	Created_at string `json:"created_at,omitempty"` // Datetime of when the cadence was created
	Creator EmbeddedResource `json:"creator,omitempty"`
	Shared bool `json:"shared,omitempty"` // Whether this cadence is visible to team members (shared)
}

// CrmUser represents the CrmUser schema from the OpenAPI specification
type CrmUser struct {
	Id int `json:"id,omitempty"` // Crm User ID
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the crm user was last updated
	User EmbeddedResource `json:"user,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the crm user was created
	Crm_id string `json:"crm_id,omitempty"` // CRM ID
}

// RecordingSetting represents the RecordingSetting schema from the OpenAPI specification
type RecordingSetting struct {
	Recording_default bool `json:"recording_default,omitempty"` // Whether this phone number should record by default
}

// EmailTemplateCounts represents the EmailTemplateCounts schema from the OpenAPI specification
type EmailTemplateCounts struct {
	Bounces int `json:"bounces,omitempty"` // The number of bounces the email template received
	Clicks int `json:"clicks,omitempty"` // The number of times links in the email template were clicked
	Replies int `json:"replies,omitempty"` // The number of replies the email template received
	Sent_emails int `json:"sent_emails,omitempty"` // The number of times the email template was sent out
	Views int `json:"views,omitempty"` // The number of times the email template was opened
}

// SuccessCounts represents the SuccessCounts schema from the OpenAPI specification
type SuccessCounts struct {
	Total_other_touches int `json:"total_other_touches,omitempty"` // The total number of other touches made in this success window
	Total_calls int `json:"total_calls,omitempty"` // The total number of calls made in this success window
	Total_emails int `json:"total_emails,omitempty"` // The total number of emails made in this success window
}

// EmbeddedAttendeeResource represents the EmbeddedAttendeeResource schema from the OpenAPI specification
type EmbeddedAttendeeResource struct {
	Status_changed bool `json:"status_changed,omitempty"` // Whether the attendee changed response status
	Deleted_at string `json:"deleted_at,omitempty"` // Datetime of when the attendee was deleted
	Email string `json:"email,omitempty"` // Email of the attendee
	Name string `json:"name,omitempty"` // Name of the attendee
	Organizer bool `json:"organizer,omitempty"` // Whether the attendee is the organizer of the event.
	Status string `json:"status,omitempty"` // The attendee's response status. Possible values are: needsAction, accepted, tentative, declined
}

// CallDisposition represents the CallDisposition schema from the OpenAPI specification
type CallDisposition struct {
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the call disposition was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the call disposition was created
	Id int `json:"id,omitempty"` // ID of CallDisposition
	Name string `json:"name,omitempty"` // An available call disposition text
}

// EmbeddedAccountCounts represents the EmbeddedAccountCounts schema from the OpenAPI specification
type EmbeddedAccountCounts struct {
	People int `json:"people,omitempty"` // Number of people in SalesLoft associated with this Account
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email_template EmbeddedResource `json:"email_template,omitempty"`
	Recipient_email_address string `json:"recipient_email_address,omitempty"` // Email address of the recipient
	User EmbeddedResource `json:"user,omitempty"`
	Click_tracking bool `json:"click_tracking,omitempty"` // Whether this email had click tracking enabled
	Status string `json:"status,omitempty"` // Status of this email through the sending process. Possible values are: sent, sent_from_gmail, sent_from_external, pending, pending_reply_check, scheduled, sending, delivering, failed, cancelled, pending_through_gmail, pending_through_external
	Subject string `json:"subject,omitempty"` // Subject of the email. This field has been determined sensitive and requires a specific scope to access it.
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the email was last updated
	Crm_activity EmbeddedResource `json:"crm_activity,omitempty"`
	Counts EmailCounts `json:"counts,omitempty"`
	Personalization string `json:"personalization,omitempty"` // Percentage of this email that has been personalized
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the email was created
	Sent_at string `json:"sent_at,omitempty"` // When this email was sent, or null if it was not sent
	Task EmbeddedResource `json:"task,omitempty"`
	View_tracking bool `json:"view_tracking,omitempty"` // Whether this email had view tracking enabled
	Action EmbeddedResource `json:"action,omitempty"`
	Error_message string `json:"error_message,omitempty"` // Error message of the email. This field has been determined sensitive and requires a specific scope to access it.
	Mailing EmbeddedResource `json:"mailing,omitempty"`
	Recipient EmbeddedResource `json:"recipient,omitempty"`
	Send_after string `json:"send_after,omitempty"` // When this email will be sent, or null if already sent
	Step EmbeddedResource `json:"step,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"` // Selected headers that are included if this email used them. Available keys are: cc, bcc
	Id int `json:"id,omitempty"` // ID of Email
	Bounced bool `json:"bounced,omitempty"` // Whether this email bounced
}

// EmailTemplateAttachment represents the EmailTemplateAttachment schema from the OpenAPI specification
type EmailTemplateAttachment struct {
	Email_template EmbeddedResource `json:"email_template,omitempty"`
	Name string `json:"name,omitempty"` // Name of the attachment
	Scanned bool `json:"scanned,omitempty"` // Checks if attachment has been scanned
	Attachment_fingerprint int `json:"attachment_fingerprint,omitempty"` // Unique attachment Identifier
	Attachment_id int `json:"attachment_id,omitempty"` // ID of the email template attachment
	Id int `json:"id,omitempty"` // ID of email template attachment association
	Attachment_content_type string `json:"attachment_content_type,omitempty"` // Content type of the attachment
	Attachment_file_size int `json:"attachment_file_size,omitempty"` // The size of the attachment
	Download_url string `json:"download_url,omitempty"` // Download url of the attachment
}

// Meeting represents the Meeting schema from the OpenAPI specification
type Meeting struct {
	Person EmbeddedResource `json:"person,omitempty"`
	Id int `json:"id,omitempty"` // ID of the meeting
	Step EmbeddedResource `json:"step,omitempty"`
	Attendees []EmbeddedAttendeeResource `json:"attendees,omitempty"` // The attendees of the meeting. Each attendee includes the following fields: status, email, name, organizer
	I_cal_uid string `json:"i_cal_uid,omitempty"` // UID of the meeting provided by target calendar provider
	Owned_by_meetings_settings EventMeetingSetting `json:"owned_by_meetings_settings,omitempty"`
	Reschedule_status string `json:"reschedule_status,omitempty"` // Status of the meeting rescheduling progress. Possible values are: pending, booked, failed, retry
	End_time string `json:"end_time,omitempty"` // End time of the meeting
	Status string `json:"status,omitempty"` // Status of the meeting. Possible values are: pending, booked, failed, retry
	Crm_references map[string]interface{} `json:"crm_references,omitempty"` // List of crm references associated with the meeting
	Recipient_email string `json:"recipient_email,omitempty"` // Email of the meeting invite recipient
	Account_id string `json:"account_id,omitempty"` // ID of the account the recipient associated to
	Event_source string `json:"event_source,omitempty"` // Source of the meeting. Possible values are: 'external' - The event was synced to Salesloft platform via Calendar Sync, 'internal' - The event was created via Salesloft platform
	All_day bool `json:"all_day,omitempty"` // Whether the meeting is an all-day meeting
	Booked_by_user EmbeddedResource `json:"booked_by_user,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the meeting was created
	Title string `json:"title,omitempty"` // Title of the meeting
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the meeting was last updated
	Calendar_type string `json:"calendar_type,omitempty"` // Calendar type of the meeting owner. Possible values are: gmail, azure, nylas, linkedin_azure, cerebro, external
	Location string `json:"location,omitempty"` // Location of the meeting
	Task_id string `json:"task_id,omitempty"` // ID of the created task
	Calendar_id string `json:"calendar_id,omitempty"` // Calendar ID of the meeting owner
	Crm_custom_fields map[string]interface{} `json:"crm_custom_fields,omitempty"` // List of crm custom fields which will be logged to SFDC
	Guests []string `json:"guests,omitempty"` // The list of attendees emails of the meeting
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	No_show bool `json:"no_show,omitempty"` // Whether the meeting is a No Show meeting
	Start_time string `json:"start_time,omitempty"` // Start time of the meeting
	Canceled_at string `json:"canceled_at,omitempty"` // Datetime of when the meeting was canceled
	Strict_attribution bool `json:"strict_attribution,omitempty"` // Strict attribution means that we 100% sure which cadence generate the meeting
	Description string `json:"description,omitempty"` // Description of the meeting
	Recipient_name string `json:"recipient_name,omitempty"` // Name of the meeting invite recipient
	Booked_by_meetings_settings EventMeetingSetting `json:"booked_by_meetings_settings,omitempty"`
	Event_id string `json:"event_id,omitempty"` // ID of the meeting created by target calendar
	Meeting_type string `json:"meeting_type,omitempty"` // Meeting type
}

// Subscription represents the Subscription schema from the OpenAPI specification
type Subscription struct {
	Enabled bool `json:"enabled,omitempty"` // Is the Webhook Subscription enabled or not
	Event_type string `json:"event_type,omitempty"` // Type of event the subscription is for
	Id int `json:"id,omitempty"` // ID for the Webhook Subscription
	Tenant_id int `json:"tenant_id,omitempty"` // ID for the tenant to which user is assigned
	User_guid string `json:"user_guid,omitempty"` // UUID of the user the token is associated with
	Callback_token string `json:"callback_token,omitempty"` // SalesLoft will include this token in the webhook event payload when calling your callback_url. It is strongly encouraged for your handler to verify this value in order to ensure the request came from SalesLoft.
	Callback_url string `json:"callback_url,omitempty"` // URL for your callback handler
}

// EmailTemplate represents the EmailTemplate schema from the OpenAPI specification
type EmailTemplate struct {
	Body string `json:"body,omitempty"` // Sanitized body of the email template without email signature
	Cadence_template bool `json:"cadence_template,omitempty"` // Whether this email template is only used on a cadence step. These templates are not visible in the SalesLoft application template list. If false, this email template is visible in the SalesLoft application, and may be used when composing an email or creating a cadence step.
	Created_at string `json:"created_at,omitempty"` // Datetime of when the email template was created
	Click_tracking_enabled bool `json:"click_tracking_enabled,omitempty"` // Whether click tracking is enabled for this email template
	Template_owner EmbeddedResource `json:"template_owner,omitempty"`
	Title string `json:"title,omitempty"` // Title of the email template
	Open_tracking_enabled bool `json:"open_tracking_enabled,omitempty"` // Whether open tracking is enabled for this email template
	Groups []EmbeddedResource `json:"groups,omitempty"` // Groups to which this template is assigned, if any
	Team_template EmbeddedResource `json:"team_template,omitempty"`
	Archived_at string `json:"archived_at,omitempty"` // Datetime of when the email template was archived, if archived
	Links map[string]interface{} `json:"_links,omitempty"` // Links to attachments and tags resources for this email template.
	Shared bool `json:"shared,omitempty"` // Whether this email template is visible to team members (shared)
	Last_used_at string `json:"last_used_at,omitempty"` // Datetime of when the email template was last used
	Subject string `json:"subject,omitempty"` // Subject of the email template
	Tags []string `json:"tags,omitempty"` // All tags applied to this email template
	Body_preview string `json:"body_preview,omitempty"` // A plain text version of the first 100 characters of the body of the email template
	Id int `json:"id,omitempty"` // ID of email template
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the email template was last updated
	Counts EmailTemplateCounts `json:"counts,omitempty"`
}

// CadenceMembership represents the CadenceMembership schema from the OpenAPI specification
type CadenceMembership struct {
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the record was last updated
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	Counts CadenceMembershipCounts `json:"counts,omitempty"`
	Currently_on_cadence bool `json:"currently_on_cadence,omitempty"` // Whether the person is currently on the cadence
	Id int `json:"id,omitempty"` // Cadence membership ID
	Person EmbeddedResource `json:"person,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the person was first added to this cadence
	Current_state string `json:"current_state,omitempty"` // The current state of the person on the cadence. Possible values are: processing: The person is being processed on a cadence. Cadence-related changes cannot be made at this time staged: The person is waiting for the first step in the cadence to occur active: The cadence has begun processing this person and is still in the process, but idle scheduled: The cadence has begun processing this person and is still in the process, with an activity scheduled to occur completed: The cadence has been completed for this person removed: The person was manually or automatically removed from the cadence removed_no_action: The person was removed from the cadence before any action occurred reassigned: The person's cadence execution was transferred to a different user, ending this user's interaction
	Person_deleted bool `json:"person_deleted,omitempty"` // Whether the associated person has since been deleted
	Latest_action EmbeddedResource `json:"latest_action,omitempty"`
	User EmbeddedResource `json:"user,omitempty"`
	Added_at string `json:"added_at,omitempty"` // Datetime of when the person was last added to this cadence
}

// MeetingUrl represents the MeetingUrl schema from the OpenAPI specification
type MeetingUrl struct {
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when MeetingUrl was last updated
	Url string `json:"url,omitempty"` // Full url of the meeting
	Created_at string `json:"created_at,omitempty"` // Datetime of when MeetingUrl was created
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Id int `json:"id,omitempty"` // ID of Tag
	Name string `json:"name,omitempty"` // Name of the tag
}

// EmbeddedRecordingResource represents the EmbeddedRecordingResource schema from the OpenAPI specification
type EmbeddedRecordingResource struct {
	Status string `json:"status,omitempty"` // The status of the call that produced this recording. Possible values are (but not limited to): no-answer: The call was not answered failed: The call was not able to be placed busy: The call was busy ringing: The call is ringing in-progress: The call is ongoing completed: The call is finished
	Url string `json:"url,omitempty"` // The url of the recording
	Recording_status string `json:"recording_status,omitempty"` // The processing status of the recording. Possible values are (but not limited to): not_recorded: there is no recording available, and there will not be one becoming available pending: the recording is currently being processed by the system processing: the recording is currently being processed by the system completed: the recording processing has been completed
}

// MimeEmailPayload represents the MimeEmailPayload schema from the OpenAPI specification
type MimeEmailPayload struct {
	Id int `json:"id,omitempty"` // Email ID
	Mailbox string `json:"mailbox,omitempty"` // Email Address of Sender's mailbox
	Message_id string `json:"message_id,omitempty"` // Unique Message ID
	Raw string `json:"raw,omitempty"` // Base64 encoded MIME email content
}

// BulkJob represents the BulkJob schema from the OpenAPI specification
type BulkJob struct {
	Processed int `json:"processed,omitempty"` // Number of processed records at the time of request for this Bulk Job
	Total int `json:"total,omitempty"` // Number of total records for this Bulk Job
	Updated_at string `json:"updated_at,omitempty"` // When this bulk job was updated
	Created_at string `json:"created_at,omitempty"` // When this bulk job was created
	Id int `json:"id,omitempty"` // ID of this Bulk Job
	Ready_to_execute bool `json:"ready_to_execute,omitempty"` // Whether the Bulk Job is ready to be executed
	Name string `json:"name,omitempty"` // Name of this Bulk Job
	Scopes []interface{} `json:"scopes,omitempty"` // Scopes
	State string `json:"state,omitempty"` // State of the Bulk Job. Must be one of: open, executing, done.
	Errors int `json:"errors,omitempty"` // Number of errored records at the time of request for this Bulk Job
	Finished_at string `json:"finished_at,omitempty"` // When this bulk job finished processing
	Marked_ready_at string `json:"marked_ready_at,omitempty"` // When this bulk job was marked as ready to execute
	Started_at string `json:"started_at,omitempty"` // When this bulk job started processing. null until bulk job is done
	TypeField string `json:"type,omitempty"` // Type of the Bulk Job.
}

// CadenceExport represents the CadenceExport schema from the OpenAPI specification
type CadenceExport struct {
	Cadence_content map[string]interface{} `json:"cadence_content,omitempty"` // The content of the cadence
}

// PendingEmail represents the PendingEmail schema from the OpenAPI specification
type PendingEmail struct {
	Mailbox string `json:"mailbox,omitempty"` // Email Address of the pending email
	Mime_email_payload EmbeddedResource `json:"mime_email_payload,omitempty"`
	Id int `json:"id,omitempty"` // ID of the email
}

// EmbeddedResource represents the EmbeddedResource schema from the OpenAPI specification
type EmbeddedResource struct {
	Id int `json:"id,omitempty"` // ID of the resource
	Href string `json:"_href,omitempty"` // Resource URL, pointed at your API version, present if this resource is available in the API
}

// TeamTemplateAttachment represents the TeamTemplateAttachment schema from the OpenAPI specification
type TeamTemplateAttachment struct {
	Team_template EmbeddedResource `json:"team_template,omitempty"`
	Attachment_file_size int `json:"attachment_file_size,omitempty"` // The size of the attachment
	Attachment_id int `json:"attachment_id,omitempty"` // ID of the team template attachment
	Download_url string `json:"download_url,omitempty"` // Download url of the attachment
	Id int `json:"id,omitempty"` // ID of team template attachment association
	Name string `json:"name,omitempty"` // Name of the attachment
}

// TeamTemplate represents the TeamTemplate schema from the OpenAPI specification
type TeamTemplate struct {
	Body_preview string `json:"body_preview,omitempty"` // A plain text version of the first 100 characters of the body of the team template
	Open_tracking_enabled bool `json:"open_tracking_enabled,omitempty"` // Whether open tracking is enabled for this team template
	Counts TeamTemplateCounts `json:"counts,omitempty"`
	Title string `json:"title,omitempty"` // Title of the team template
	Last_modified_user EmbeddedResource `json:"last_modified_user,omitempty"`
	Tags []string `json:"tags,omitempty"` // All tags applied to this team template
	Id string `json:"id,omitempty"` // ID of team template
	Subject string `json:"subject,omitempty"` // Subject of the team template
	Archived_at string `json:"archived_at,omitempty"` // Datetime of when the team template was archived, if archived
	Body string `json:"body,omitempty"` // Body of the team template
	Last_used_at string `json:"last_used_at,omitempty"` // Datetime of when the team template was last used
	Links map[string]interface{} `json:"_links,omitempty"` // Links to attachments resource for this template
	Created_at string `json:"created_at,omitempty"` // Datetime of when the team template was created
	Last_modified_at string `json:"last_modified_at,omitempty"` // Datetime of when the team template was last modified
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the team template was last updated
	Click_tracking_enabled bool `json:"click_tracking_enabled,omitempty"` // Whether click tracking is enabled for this team template
}

// JobDataCreationResult represents the JobDataCreationResult schema from the OpenAPI specification
type JobDataCreationResult struct {
	Records int `json:"records,omitempty"` // Number of records created
}

// AccountTier represents the AccountTier schema from the OpenAPI specification
type AccountTier struct {
	Created_at string `json:"created_at,omitempty"` // Datetime of when the Account Tier was created
	Id int `json:"id,omitempty"` // ID of Account Tier
	Name string `json:"name,omitempty"` // Name of the Account Tier
	Order int `json:"order,omitempty"` // The order of the account tier
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the Account Tier was last updated
}

// Call represents the Call schema from the OpenAPI specification
type Call struct {
	Disposition string `json:"disposition,omitempty"` // Result of the call
	To string `json:"to,omitempty"` // Phone number that received the call
	Duration int `json:"duration,omitempty"` // Length of the call in seconds
	Sentiment string `json:"sentiment,omitempty"` // Outcome of the conversation
	Called_person EmbeddedResource `json:"called_person,omitempty"`
	Cadence EmbeddedResource `json:"cadence,omitempty"`
	User EmbeddedResource `json:"user,omitempty"`
	Id int `json:"id,omitempty"` // ID of Call
	Note EmbeddedResource `json:"note,omitempty"`
	Recordings []EmbeddedRecordingResource `json:"recordings,omitempty"` // The recordings for this this call and their status
	Created_at string `json:"created_at,omitempty"` // Datetime of when the call was created
	Crm_activity EmbeddedResource `json:"crm_activity,omitempty"`
	Step EmbeddedResource `json:"step,omitempty"`
	Action EmbeddedResource `json:"action,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the call was last updated
}

// Success represents the Success schema from the OpenAPI specification
type Success struct {
	Counts SuccessCounts `json:"counts,omitempty"`
	Id int `json:"id,omitempty"` // ID of success
	Latest_email EmbeddedResource `json:"latest_email,omitempty"`
	Person EmbeddedResource `json:"person,omitempty"`
	Success_window_started_at string `json:"success_window_started_at,omitempty"` // Datetime of when this person was first worked, leading up to the success
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the success was last updated
	User EmbeddedResource `json:"user,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Datetime of when the success was created
	Latest_action EmbeddedResource `json:"latest_action,omitempty"`
	Latest_cadence EmbeddedResource `json:"latest_cadence,omitempty"`
	Latest_call EmbeddedResource `json:"latest_call,omitempty"`
	Latest_step EmbeddedResource `json:"latest_step,omitempty"`
	Succeeded_at string `json:"succeeded_at,omitempty"` // Datetime of when the success was recorded
}

// CrmActivity represents the CrmActivity schema from the OpenAPI specification
type CrmActivity struct {
	Created_at string `json:"created_at,omitempty"` // Datetime of when the crm activity was created
	Person EmbeddedResource `json:"person,omitempty"`
	Crm_id string `json:"crm_id,omitempty"` // The ID of the activity in your CRM, if written to your CRM
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the crm activity was last updated
	Subject string `json:"subject,omitempty"` // The subject field of the activity in your CRM
	User EmbeddedResource `json:"user,omitempty"`
	Activity_type string `json:"activity_type,omitempty"` // The type of activity that is being recorded, if available. The values can change over time, but could be one of: email, phone, email reminder, inmail
	Custom_crm_fields map[string]interface{} `json:"custom_crm_fields,omitempty"` // Additional fields that are logged to your CRM, if mapped by the team at the time of writing to your CRM
	Description string `json:"description,omitempty"` // The description field of the activity in your CRM
	ErrorField string `json:"error,omitempty"` // Information about why this crm activity failed to sync, if it did fail to sync. Failed activities will be automatically retried and may become successful in the future
	Id int `json:"id,omitempty"` // CrmActivity ID
}

// BulkJobResult represents the BulkJobResult schema from the OpenAPI specification
type BulkJobResult struct {
	Status string `json:"status,omitempty"` // Status of the record that was processed. Will be one of: success, error
	ErrorField string `json:"error,omitempty"` // Error message for the record that was processed. Will be null if there was no error.
	Id int `json:"id,omitempty"` // ID of the record that was processed
	Record map[string]interface{} `json:"record,omitempty"` // The data that was used to process the operation
	Resource map[string]interface{} `json:"resource,omitempty"` // The object containing the resulting resource from performing the bulk action on this record
}

// ExternalEmail represents the ExternalEmail schema from the OpenAPI specification
type ExternalEmail struct {
	Message_id string `json:"message_id,omitempty"` // Message id present in the External Email header
}

// PhoneNumberAssignment represents the PhoneNumberAssignment schema from the OpenAPI specification
type PhoneNumberAssignment struct {
	Id int `json:"id,omitempty"` // PhoneNumberAssignment ID
	Number string `json:"number,omitempty"` // The phone number associated with this assignment
	User EmbeddedResource `json:"user,omitempty"`
}

// PersonUpsert represents the PersonUpsert schema from the OpenAPI specification
type PersonUpsert struct {
	Upsert_type string `json:"upsert_type,omitempty"` // The type of upsert. One of: create, update
	Person Person `json:"person,omitempty"`
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Email_client_configured bool `json:"email_client_configured,omitempty"` // Whether this user has a email client configured
	Slack_username string `json:"slack_username,omitempty"` // Slack username
	Work_country string `json:"work_country,omitempty"` // Work Country
	Last_name string `json:"last_name,omitempty"` // Last name of user
	Email_signature_type string `json:"email_signature_type,omitempty"` // Email signature type
	Team_admin bool `json:"team_admin,omitempty"` // Team Admin
	External_feature_flags map[string]interface{} `json:"external_feature_flags,omitempty"` // Feature flags that are for this user. New flags may appear or disappear at any time
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the user was last updated
	Created_at string `json:"created_at,omitempty"` // Datetime of when the user was created
	Local_dial_enabled bool `json:"local_dial_enabled,omitempty"` // Whether this user has Local Dial enabled
	Name string `json:"name,omitempty"` // Display name of user
	Sending_email_address string `json:"sending_email_address,omitempty"` // The email address that email of the user will be sent from, resolved in the following resolution order: from_user, email_client_email_address, email
	First_name string `json:"first_name,omitempty"` // First name of user
	Phone_number_assignment EmbeddedResource `json:"phone_number_assignment,omitempty"`
	Bcc_email_address string `json:"bcc_email_address,omitempty"` // Address that will be BBC'd on all emails from this user
	Guid string `json:"guid,omitempty"` // Globally unique user ID. New endpoints will explicitly accept this over id
	Role EmbeddedResource `json:"role,omitempty"`
	Email_signature string `json:"email_signature,omitempty"` // Email signature
	Click_to_call_enabled bool `json:"click_to_call_enabled,omitempty"` // Whether this user has click to call enabled
	Group EmbeddedResource `json:"group,omitempty"`
	Id int `json:"id,omitempty"` // User ID
	Private_fields map[string]interface{} `json:"_private_fields,omitempty"` // For internal use only. This field does not comply with our backwards compatability policies.
	Email_client_email_address string `json:"email_client_email_address,omitempty"` // Email address associated with the email client of the user
	Crm_connected bool `json:"crm_connected,omitempty"` // Whether the user has a crm connected
	Team EmbeddedResource `json:"team,omitempty"`
	Phone_client EmbeddedResource `json:"phone_client,omitempty"`
	Email string `json:"email,omitempty"` // Email address provided to accounts.salesloft.com
	Full_email_address string `json:"full_email_address,omitempty"` // RFC 5322 compliant email address
	Email_signature_click_tracking_disabled bool `json:"email_signature_click_tracking_disabled,omitempty"` // Whether this user has click tracking disabled in email signature
	Active bool `json:"active,omitempty"` // Whether an user is currently active in SalesLoft
	Job_role string `json:"job_role,omitempty"` // Job role of user
	From_address string `json:"from_address,omitempty"` // The from address of this user
	Time_zone string `json:"time_zone,omitempty"` // User Time Zone
	Twitter_handle string `json:"twitter_handle,omitempty"` // Twitter handle
}

// LiveWebsiteTrackingParameter represents the LiveWebsiteTrackingParameter schema from the OpenAPI specification
type LiveWebsiteTrackingParameter struct {
	Parameters []map[string]interface{} `json:"parameters,omitempty"` // A SalesLoft identifier
}

// CadenceMembershipCounts represents the CadenceMembershipCounts schema from the OpenAPI specification
type CadenceMembershipCounts struct {
	Bounces int `json:"bounces,omitempty"` // The number of times emails sent from the cadence to the person bounced
	Calls int `json:"calls,omitempty"` // The number of times a call was logged from the cadence to the person
	Clicks int `json:"clicks,omitempty"` // The number of times emails sent from the cadence to the person were clicked
	Replies int `json:"replies,omitempty"` // The number of times emails sent from the cadence to the person were replied to
	Sent_emails int `json:"sent_emails,omitempty"` // The number of times emails were sent from the cadence to the person
	Views int `json:"views,omitempty"` // The number of times emails sent from the cadence to the person were opened
}

// TeamTemplateCounts represents the TeamTemplateCounts schema from the OpenAPI specification
type TeamTemplateCounts struct {
	Bounces int `json:"bounces,omitempty"` // The number of bounces the team template received
	Clicks int `json:"clicks,omitempty"` // The number of times links in the team template were clicked
	Replies int `json:"replies,omitempty"` // The number of replies the team template received
	Sent_emails int `json:"sent_emails,omitempty"` // The number of times the team template was sent out
	Views int `json:"views,omitempty"` // The number of times the team template was opened
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Parent_id int `json:"parent_id,omitempty"` // ID of the parent Group
	Accessible_groups []EmbeddedResource `json:"accessible_groups,omitempty"` // Groups accessible if any
	Id int `json:"id,omitempty"` // ID of the Group
	Name string `json:"name,omitempty"` // Name of the Group
}

// CallDataRecord represents the CallDataRecord schema from the OpenAPI specification
type CallDataRecord struct {
	Status string `json:"status,omitempty"` // The outcome of the call. Can be one of: queued, initiated, ringing, in-progress, completed, busy, no-answer, canceled, failed
	User EmbeddedResource `json:"user,omitempty"`
	Called_person EmbeddedResource `json:"called_person,omitempty"`
	Direction string `json:"direction,omitempty"` // Direction of the call. Can be one of: inbound, outbound
	Recording EmbeddedRecordingResource `json:"recording,omitempty"`
	Call EmbeddedResource `json:"call,omitempty"`
	Call_uuid string `json:"call_uuid,omitempty"` // UUID of the call. Legs of the same call will have the same call_uuid.
	Id int `json:"id,omitempty"` // ID of CallDataRecord
	Updated_at string `json:"updated_at,omitempty"` // Datetime of when the call was last updated
	Call_type string `json:"call_type,omitempty"` // Type of the call. Can be one of: call, bridge, collaboration. Though exact values may change over time
	Created_at string `json:"created_at,omitempty"` // Datetime of when the call was created
	To string `json:"to,omitempty"` // Phone number that received the call
	Duration int `json:"duration,omitempty"` // Length of the call in seconds
	From string `json:"from,omitempty"` // Phone number that placed the call
}

// CadenceCounts represents the CadenceCounts schema from the OpenAPI specification
type CadenceCounts struct {
	Target_daily_people int `json:"target_daily_people,omitempty"` // The user defined target for number of people to add to the cadence each day
	Cadence_people int `json:"cadence_people,omitempty"` // The number of people that have ever been added to the cadence
	Meetings_booked int `json:"meetings_booked,omitempty"` // The number of meetings booked and attributed to the cadence
	Opportunities_created int `json:"opportunities_created,omitempty"` // The number of opportunities created and attributed to the cadence
	People_acted_on_count int `json:"people_acted_on_count,omitempty"` // The number of people that have been skipped, scheduled, or advanced in a cadence
}
