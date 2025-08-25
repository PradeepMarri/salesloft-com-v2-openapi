package main

import (
	"github.com/salesloft-platform/mcp-server/config"
	"github.com/salesloft-platform/mcp-server/models"
	tools_email_templates "github.com/salesloft-platform/mcp-server/tools/email_templates"
	tools_team_template_attachments "github.com/salesloft-platform/mcp-server/tools/team_template_attachments"
	tools_people "github.com/salesloft-platform/mcp-server/tools/people"
	tools_imports "github.com/salesloft-platform/mcp-server/tools/imports"
	tools_call_data_records "github.com/salesloft-platform/mcp-server/tools/call_data_records"
	tools_account_stages "github.com/salesloft-platform/mcp-server/tools/account_stages"
	tools_crm_activities "github.com/salesloft-platform/mcp-server/tools/crm_activities"
	tools_notes "github.com/salesloft-platform/mcp-server/tools/notes"
	tools_calls "github.com/salesloft-platform/mcp-server/tools/calls"
	tools_me "github.com/salesloft-platform/mcp-server/tools/me"
	tools_cadences "github.com/salesloft-platform/mcp-server/tools/cadences"
	tools_cadence_memberships "github.com/salesloft-platform/mcp-server/tools/cadence_memberships"
	tools_accounts "github.com/salesloft-platform/mcp-server/tools/accounts"
	tools_action_details_call_instructions "github.com/salesloft-platform/mcp-server/tools/action_details_call_instructions"
	tools_email_template_attachments "github.com/salesloft-platform/mcp-server/tools/email_template_attachments"
	tools_caller_ids "github.com/salesloft-platform/mcp-server/tools/caller_ids"
	tools_emails "github.com/salesloft-platform/mcp-server/tools/emails"
	tools_roles "github.com/salesloft-platform/mcp-server/tools/roles"
	tools_account_tiers "github.com/salesloft-platform/mcp-server/tools/account_tiers"
	tools_bulk_jobs_job_data "github.com/salesloft-platform/mcp-server/tools/bulk_jobs_job_data"
	tools_phone_number_assignments "github.com/salesloft-platform/mcp-server/tools/phone_number_assignments"
	tools_recording_settings "github.com/salesloft-platform/mcp-server/tools/recording_settings"
	tools_custom_fields "github.com/salesloft-platform/mcp-server/tools/custom_fields"
	tools_webhook_subscriptions "github.com/salesloft-platform/mcp-server/tools/webhook_subscriptions"
	tools_meetings "github.com/salesloft-platform/mcp-server/tools/meetings"
	tools_bulk_jobs "github.com/salesloft-platform/mcp-server/tools/bulk_jobs"
	tools_call_sentiments "github.com/salesloft-platform/mcp-server/tools/call_sentiments"
	tools_crm_activity_fields "github.com/salesloft-platform/mcp-server/tools/crm_activity_fields"
	tools_actions "github.com/salesloft-platform/mcp-server/tools/actions"
	tools_team_templates "github.com/salesloft-platform/mcp-server/tools/team_templates"
	tools_tags "github.com/salesloft-platform/mcp-server/tools/tags"
	tools_groups "github.com/salesloft-platform/mcp-server/tools/groups"
	tools_tasks "github.com/salesloft-platform/mcp-server/tools/tasks"
	tools_steps "github.com/salesloft-platform/mcp-server/tools/steps"
	tools_activity_histories "github.com/salesloft-platform/mcp-server/tools/activity_histories"
	tools_meetings_settings_searches "github.com/salesloft-platform/mcp-server/tools/meetings_settings_searches"
	tools_bulk_jobs_results "github.com/salesloft-platform/mcp-server/tools/bulk_jobs_results"
	tools_saved_list_views "github.com/salesloft-platform/mcp-server/tools/saved_list_views"
	tools_person_stages "github.com/salesloft-platform/mcp-server/tools/person_stages"
	tools_crm_users "github.com/salesloft-platform/mcp-server/tools/crm_users"
	tools_pending_emails "github.com/salesloft-platform/mcp-server/tools/pending_emails"
	tools_users "github.com/salesloft-platform/mcp-server/tools/users"
	tools_calendar_events "github.com/salesloft-platform/mcp-server/tools/calendar_events"
	tools_successes "github.com/salesloft-platform/mcp-server/tools/successes"
	tools_team "github.com/salesloft-platform/mcp-server/tools/team"
	tools_call_dispositions "github.com/salesloft-platform/mcp-server/tools/call_dispositions"
	tools_cadence_exports "github.com/salesloft-platform/mcp-server/tools/cadence_exports"
	tools_mime_email_payloads "github.com/salesloft-platform/mcp-server/tools/mime_email_payloads"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_email_templates.CreateGet_v2_email_templates_jsonTool(cfg),
		tools_team_template_attachments.CreateGet_v2_team_template_attachments_jsonTool(cfg),
		tools_people.CreateGet_v2_people_jsonTool(cfg),
		tools_imports.CreateGet_v2_imports_id_jsonTool(cfg),
		tools_imports.CreateDelete_v2_imports_id_jsonTool(cfg),
		tools_call_data_records.CreateGet_v2_call_data_records_jsonTool(cfg),
		tools_account_stages.CreateGet_v2_account_stages_jsonTool(cfg),
		tools_crm_activities.CreateGet_v2_crm_activities_id_jsonTool(cfg),
		tools_notes.CreateDelete_v2_notes_id_jsonTool(cfg),
		tools_notes.CreateGet_v2_notes_id_jsonTool(cfg),
		tools_calls.CreateGet_v2_activities_calls_jsonTool(cfg),
		tools_me.CreateGet_v2_me_jsonTool(cfg),
		tools_cadences.CreateGet_v2_cadences_jsonTool(cfg),
		tools_cadence_memberships.CreateGet_v2_cadence_memberships_jsonTool(cfg),
		tools_cadence_memberships.CreatePost_v2_cadence_memberships_jsonTool(cfg),
		tools_accounts.CreateGet_v2_accounts_jsonTool(cfg),
		tools_action_details_call_instructions.CreateGet_v2_action_details_call_instructions_id_jsonTool(cfg),
		tools_email_template_attachments.CreateGet_v2_email_template_attachments_jsonTool(cfg),
		tools_caller_ids.CreateGet_v2_phone_numbers_caller_ids_jsonTool(cfg),
		tools_emails.CreateGet_v2_activities_emails_id_jsonTool(cfg),
		tools_roles.CreateGet_v2_custom_roles_jsonTool(cfg),
		tools_notes.CreateGet_v2_notes_jsonTool(cfg),
		tools_account_tiers.CreateGet_v2_account_tiers_jsonTool(cfg),
		tools_bulk_jobs_job_data.CreateGet_v2_bulk_jobs_bulk_jobs_id_job_dataTool(cfg),
		tools_phone_number_assignments.CreateGet_v2_phone_number_assignments_jsonTool(cfg),
		tools_recording_settings.CreateGet_v2_phone_numbers_recording_settings_id_jsonTool(cfg),
		tools_custom_fields.CreateGet_v2_custom_fields_jsonTool(cfg),
		tools_webhook_subscriptions.CreateDelete_v2_webhook_subscriptions_idTool(cfg),
		tools_webhook_subscriptions.CreateGet_v2_webhook_subscriptions_idTool(cfg),
		tools_meetings.CreateGet_v2_meetings_jsonTool(cfg),
		tools_bulk_jobs.CreateGet_v2_bulk_jobsTool(cfg),
		tools_call_sentiments.CreateGet_v2_call_sentiments_jsonTool(cfg),
		tools_crm_activities.CreateGet_v2_crm_activities_jsonTool(cfg),
		tools_crm_activity_fields.CreateGet_v2_crm_activity_fields_jsonTool(cfg),
		tools_actions.CreateGet_v2_actions_id_jsonTool(cfg),
		tools_emails.CreateGet_v2_activities_emails_jsonTool(cfg),
		tools_team_templates.CreateGet_v2_team_templates_jsonTool(cfg),
		tools_actions.CreateGet_v2_actions_jsonTool(cfg),
		tools_custom_fields.CreateGet_v2_custom_fields_id_jsonTool(cfg),
		tools_custom_fields.CreateDelete_v2_custom_fields_id_jsonTool(cfg),
		tools_tags.CreateGet_v2_tags_jsonTool(cfg),
		tools_groups.CreateGet_v2_groups_jsonTool(cfg),
		tools_tasks.CreateGet_v2_tasks_jsonTool(cfg),
		tools_groups.CreateGet_v2_groups_id_jsonTool(cfg),
		tools_account_tiers.CreateGet_v2_account_tiers_id_jsonTool(cfg),
		tools_call_data_records.CreateGet_v2_call_data_records_id_jsonTool(cfg),
		tools_calls.CreateGet_v2_activities_calls_id_jsonTool(cfg),
		tools_steps.CreateGet_v2_steps_jsonTool(cfg),
		tools_account_stages.CreateGet_v2_account_stages_id_jsonTool(cfg),
		tools_imports.CreateGet_v2_imports_jsonTool(cfg),
		tools_activity_histories.CreateGet_v2_activity_historiesTool(cfg),
		tools_email_templates.CreateGet_v2_email_templates_id_jsonTool(cfg),
		tools_meetings_settings_searches.CreatePost_v2_meetings_settings_searches_jsonTool(cfg),
		tools_webhook_subscriptions.CreateGet_v2_webhook_subscriptionsTool(cfg),
		tools_bulk_jobs_results.CreateGet_v2_bulk_jobs_bulk_jobs_id_resultsTool(cfg),
		tools_team_templates.CreateGet_v2_team_templates_id_jsonTool(cfg),
		tools_roles.CreateGet_v2_custom_roles_id_jsonTool(cfg),
		tools_saved_list_views.CreateGet_v2_saved_list_views_jsonTool(cfg),
		tools_person_stages.CreateDelete_v2_person_stages_id_jsonTool(cfg),
		tools_person_stages.CreateGet_v2_person_stages_id_jsonTool(cfg),
		tools_crm_users.CreateGet_v2_crm_users_jsonTool(cfg),
		tools_pending_emails.CreateGet_v2_pending_emails_jsonTool(cfg),
		tools_phone_number_assignments.CreateGet_v2_phone_number_assignments_id_jsonTool(cfg),
		tools_accounts.CreateDelete_v2_accounts_id_jsonTool(cfg),
		tools_accounts.CreateGet_v2_accounts_id_jsonTool(cfg),
		tools_people.CreateDelete_v2_people_id_jsonTool(cfg),
		tools_people.CreateGet_v2_people_id_jsonTool(cfg),
		tools_tasks.CreateGet_v2_tasks_id_jsonTool(cfg),
		tools_person_stages.CreateGet_v2_person_stages_jsonTool(cfg),
		tools_saved_list_views.CreateDelete_v2_saved_list_views_id_jsonTool(cfg),
		tools_saved_list_views.CreateGet_v2_saved_list_views_id_jsonTool(cfg),
		tools_steps.CreateGet_v2_steps_id_jsonTool(cfg),
		tools_users.CreateGet_v2_users_id_jsonTool(cfg),
		tools_cadences.CreateGet_v2_cadences_id_jsonTool(cfg),
		tools_users.CreateGet_v2_users_jsonTool(cfg),
		tools_bulk_jobs.CreateGet_v2_bulk_jobs_idTool(cfg),
		tools_calendar_events.CreateGet_v2_calendar_eventsTool(cfg),
		tools_successes.CreateGet_v2_successes_jsonTool(cfg),
		tools_cadence_memberships.CreateDelete_v2_cadence_memberships_id_jsonTool(cfg),
		tools_cadence_memberships.CreateGet_v2_cadence_memberships_id_jsonTool(cfg),
		tools_team.CreateGet_v2_team_jsonTool(cfg),
		tools_call_dispositions.CreateGet_v2_call_dispositions_jsonTool(cfg),
		tools_cadence_exports.CreateGet_v2_cadence_exports_id_jsonTool(cfg),
		tools_mime_email_payloads.CreateGet_v2_mime_email_payloads_id_jsonTool(cfg),
		tools_action_details_call_instructions.CreateGet_v2_action_details_call_instructions_jsonTool(cfg),
	}
}
