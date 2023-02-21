package main

import "time"

type WITCRCreate struct {
	SubscriptionID string `json:"subscriptionId"`
	NotificationID int    `json:"notificationId"`
	ID             string `json:"id"`
	EventType      string `json:"eventType"`
	PublisherID    string `json:"publisherId"`
	Message        struct {
		HTML string `json:"html"`
	} `json:"message"`
	DetailedMessage struct {
		HTML string `json:"html"`
	} `json:"detailedMessage"`
	Resource struct {
		ID     int `json:"id"`
		Rev    int `json:"rev"`
		Fields struct {
			SystemAreaPath                                      string    `json:"System.AreaPath"`
			SystemTeamProject                                   string    `json:"System.TeamProject"`
			SystemIterationPath                                 string    `json:"System.IterationPath"`
			SystemWorkItemType                                  string    `json:"System.WorkItemType"`
			SystemState                                         string    `json:"System.State"`
			SystemReason                                        string    `json:"System.Reason"`
			SystemCreatedDate                                   time.Time `json:"System.CreatedDate"`
			SystemCreatedBy                                     string    `json:"System.CreatedBy"`
			SystemChangedDate                                   time.Time `json:"System.ChangedDate"`
			SystemChangedBy                                     string    `json:"System.ChangedBy"`
			SystemCommentCount                                  int       `json:"System.CommentCount"`
			SystemTitle                                         string    `json:"System.Title"`
			SystemBoardColumn                                   string    `json:"System.BoardColumn"`
			SystemBoardColumnDone                               bool      `json:"System.BoardColumnDone"`
			MicrosoftVSTSCommonPriority                         int       `json:"Microsoft.VSTS.Common.Priority"`
			MicrosoftVSTSCommonSeverity                         string    `json:"Microsoft.VSTS.Common.Severity"`
			MicrosoftVSTSCommonValueArea                        string    `json:"Microsoft.VSTS.Common.ValueArea"`
			BmClient                                            string    `json:"bm.Client"`
			BmPlatform                                          string    `json:"bm.Platform"`
			WEF54E1C90D03E24593955A6C6D847649D5KanbanColumn     string    `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column"`
			WEF54E1C90D03E24593955A6C6D847649D5KanbanColumnDone bool      `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column.Done"`
			WEF461E61775E55407092925710D059FC79KanbanColumn     string    `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column"`
			WEF461E61775E55407092925710D059FC79KanbanColumnDone bool      `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column.Done"`
			BmZendesk                                           bool      `json:"bm.Zendesk"`
			BmToEstimate                                        bool      `json:"bm.ToEstimate"`
			BmContract                                          bool      `json:"bm.Contract"`
			TimeFunctionalRequirementsEstimate                  int       `json:"time.FunctionalRequirementsEstimate"`
			TimeUIUxEstimate                                    int       `json:"time.UiUxEstimate"`
			TimeFuncReqCompleteWork                             int       `json:"time.FuncReqCompleteWork"`
			TimeUIUxCompleted                                   int       `json:"time.UiUxCompleted"`
			SystemDescription                                   string    `json:"System.Description"`
			MicrosoftVSTSTCMSystemInfo                          string    `json:"Microsoft.VSTS.TCM.SystemInfo"`
			MicrosoftVSTSTCMReproSteps                          string    `json:"Microsoft.VSTS.TCM.ReproSteps"`
			MicrosoftVSTSCommonAcceptanceCriteria               string    `json:"Microsoft.VSTS.Common.AcceptanceCriteria"`
		} `json:"fields"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			WorkItemUpdates struct {
				Href string `json:"href"`
			} `json:"workItemUpdates"`
			WorkItemRevisions struct {
				Href string `json:"href"`
			} `json:"workItemRevisions"`
			WorkItemComments struct {
				Href string `json:"href"`
			} `json:"workItemComments"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			WorkItemType struct {
				Href string `json:"href"`
			} `json:"workItemType"`
			Fields struct {
				Href string `json:"href"`
			} `json:"fields"`
		} `json:"_links"`
		URL string `json:"url"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Server struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"server"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}

type WITCIUpdate struct {
	SubscriptionID string `json:"subscriptionId"`
	NotificationID int    `json:"notificationId"`
	ID             string `json:"id"`
	EventType      string `json:"eventType"`
	PublisherID    string `json:"publisherId"`
	Message        struct {
		HTML string `json:"html"`
	} `json:"message"`
	DetailedMessage struct {
		HTML string `json:"html"`
	} `json:"detailedMessage"`
	Resource struct {
		ID         int `json:"id"`
		WorkItemID int `json:"workItemId"`
		Rev        int `json:"rev"`
		RevisedBy  struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"revisedBy"`
		RevisedDate time.Time `json:"revisedDate"`
		Fields      struct {
			SystemRev struct {
				OldValue int `json:"oldValue"`
				NewValue int `json:"newValue"`
			} `json:"System.Rev"`
			SystemAuthorizedDate struct {
				OldValue time.Time `json:"oldValue"`
				NewValue time.Time `json:"newValue"`
			} `json:"System.AuthorizedDate"`
			SystemRevisedDate struct {
				OldValue time.Time `json:"oldValue"`
				NewValue time.Time `json:"newValue"`
			} `json:"System.RevisedDate"`
			SystemChangedDate struct {
				OldValue time.Time `json:"oldValue"`
				NewValue time.Time `json:"newValue"`
			} `json:"System.ChangedDate"`
			SystemWatermark struct {
				OldValue int `json:"oldValue"`
				NewValue int `json:"newValue"`
			} `json:"System.Watermark"`
			MicrosoftVSTSTCMSystemInfo struct {
				OldValue string `json:"oldValue"`
				NewValue string `json:"newValue"`
			} `json:"Microsoft.VSTS.TCM.SystemInfo"`
		} `json:"fields"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			WorkItemUpdates struct {
				Href string `json:"href"`
			} `json:"workItemUpdates"`
			Parent struct {
				Href string `json:"href"`
			} `json:"parent"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"_links"`
		URL      string `json:"url"`
		Revision struct {
			ID     int `json:"id"`
			Rev    int `json:"rev"`
			Fields struct {
				SystemAreaPath                                      string    `json:"System.AreaPath"`
				SystemTeamProject                                   string    `json:"System.TeamProject"`
				SystemIterationPath                                 string    `json:"System.IterationPath"`
				SystemWorkItemType                                  string    `json:"System.WorkItemType"`
				SystemState                                         string    `json:"System.State"`
				SystemReason                                        string    `json:"System.Reason"`
				SystemCreatedDate                                   time.Time `json:"System.CreatedDate"`
				SystemCreatedBy                                     string    `json:"System.CreatedBy"`
				SystemChangedDate                                   time.Time `json:"System.ChangedDate"`
				SystemChangedBy                                     string    `json:"System.ChangedBy"`
				SystemCommentCount                                  int       `json:"System.CommentCount"`
				SystemTitle                                         string    `json:"System.Title"`
				SystemBoardColumn                                   string    `json:"System.BoardColumn"`
				SystemBoardColumnDone                               bool      `json:"System.BoardColumnDone"`
				MicrosoftVSTSCommonPriority                         int       `json:"Microsoft.VSTS.Common.Priority"`
				MicrosoftVSTSCommonSeverity                         string    `json:"Microsoft.VSTS.Common.Severity"`
				MicrosoftVSTSCommonValueArea                        string    `json:"Microsoft.VSTS.Common.ValueArea"`
				BmClient                                            string    `json:"bm.Client"`
				BmPlatform                                          string    `json:"bm.Platform"`
				WEF54E1C90D03E24593955A6C6D847649D5KanbanColumn     string    `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column"`
				WEF54E1C90D03E24593955A6C6D847649D5KanbanColumnDone bool      `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column.Done"`
				BmServerVersion                                     string    `json:"bm.ServerVersion"`
				BmIpadVersion                                       string    `json:"bm.IpadVersion"`
				BmAndroidVersion                                    string    `json:"bm.AndroidVersion"`
				WEF461E61775E55407092925710D059FC79KanbanColumn     string    `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column"`
				WEF461E61775E55407092925710D059FC79KanbanColumnDone bool      `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column.Done"`
				BmZendesk                                           bool      `json:"bm.Zendesk"`
				BmContract                                          bool      `json:"bm.Contract"`
				TimeCompletedWorkQA                                 int       `json:"time.CompletedWorkQA"`
				SystemDescription                                   string    `json:"System.Description"`
				MicrosoftVSTSTCMSystemInfo                          string    `json:"Microsoft.VSTS.TCM.SystemInfo"`
			} `json:"fields"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				WorkItemRevisions struct {
					Href string `json:"href"`
				} `json:"workItemRevisions"`
				Parent struct {
					Href string `json:"href"`
				} `json:"parent"`
			} `json:"_links"`
			URL string `json:"url"`
		} `json:"revision"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Server struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"server"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}

type WITCICreate struct {
	SubscriptionID string `json:"subscriptionId"`
	NotificationID int    `json:"notificationId"`
	ID             string `json:"id"`
	EventType      string `json:"eventType"`
	PublisherID    string `json:"publisherId"`
	Message        struct {
		HTML string `json:"html"`
	} `json:"message"`
	DetailedMessage struct {
		HTML string `json:"html"`
	} `json:"detailedMessage"`
	Resource struct {
		ID     int `json:"id"`
		Rev    int `json:"rev"`
		Fields struct {
			SystemAreaPath                                      string    `json:"System.AreaPath"`
			SystemTeamProject                                   string    `json:"System.TeamProject"`
			SystemIterationPath                                 string    `json:"System.IterationPath"`
			SystemWorkItemType                                  string    `json:"System.WorkItemType"`
			SystemState                                         string    `json:"System.State"`
			SystemReason                                        string    `json:"System.Reason"`
			SystemAssignedTo                                    string    `json:"System.AssignedTo"`
			SystemCreatedDate                                   time.Time `json:"System.CreatedDate"`
			SystemCreatedBy                                     string    `json:"System.CreatedBy"`
			SystemChangedDate                                   time.Time `json:"System.ChangedDate"`
			SystemChangedBy                                     string    `json:"System.ChangedBy"`
			SystemCommentCount                                  int       `json:"System.CommentCount"`
			SystemTitle                                         string    `json:"System.Title"`
			SystemBoardColumn                                   string    `json:"System.BoardColumn"`
			SystemBoardColumnDone                               bool      `json:"System.BoardColumnDone"`
			MicrosoftVSTSCommonPriority                         int       `json:"Microsoft.VSTS.Common.Priority"`
			MicrosoftVSTSCommonSeverity                         string    `json:"Microsoft.VSTS.Common.Severity"`
			MicrosoftVSTSCommonValueArea                        string    `json:"Microsoft.VSTS.Common.ValueArea"`
			BmClient                                            string    `json:"bm.Client"`
			BmPlatform                                          string    `json:"bm.Platform"`
			WEF54E1C90D03E24593955A6C6D847649D5KanbanColumn     string    `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column"`
			WEF54E1C90D03E24593955A6C6D847649D5KanbanColumnDone bool      `json:"WEF_54E1C90D03E24593955A6C6D847649D5_Kanban.Column.Done"`
			BmServerVersion                                     string    `json:"bm.ServerVersion"`
			BmIpadVersion                                       string    `json:"bm.IpadVersion"`
			BmAndroidVersion                                    string    `json:"bm.AndroidVersion"`
			WEF461E61775E55407092925710D059FC79KanbanColumn     string    `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column"`
			WEF461E61775E55407092925710D059FC79KanbanColumnDone bool      `json:"WEF_461E61775E55407092925710D059FC79_Kanban.Column.Done"`
			BmZendesk                                           bool      `json:"bm.Zendesk"`
			BmContract                                          bool      `json:"bm.Contract"`
			TimeCompletedWorkQA                                 int       `json:"time.CompletedWorkQA"`
			SystemDescription                                   string    `json:"System.Description"`
			MicrosoftVSTSTCMSystemInfo                          string    `json:"Microsoft.VSTS.TCM.SystemInfo"`
			MicrosoftVSTSTCMReproSteps                          string    `json:"Microsoft.VSTS.TCM.ReproSteps"`
			MicrosoftVSTSCommonAcceptanceCriteria               string    `json:"Microsoft.VSTS.Common.AcceptanceCriteria"`
		} `json:"fields"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			WorkItemUpdates struct {
				Href string `json:"href"`
			} `json:"workItemUpdates"`
			WorkItemRevisions struct {
				Href string `json:"href"`
			} `json:"workItemRevisions"`
			WorkItemComments struct {
				Href string `json:"href"`
			} `json:"workItemComments"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			WorkItemType struct {
				Href string `json:"href"`
			} `json:"workItemType"`
			Fields struct {
				Href string `json:"href"`
			} `json:"fields"`
		} `json:"_links"`
		URL string `json:"url"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Server struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"server"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}

type Build struct {
	SubscriptionID string `json:"subscriptionId"`
	NotificationID int    `json:"notificationId"`
	ID             string `json:"id"`
	EventType      string `json:"eventType"`
	PublisherID    string `json:"publisherId"`
	Message        struct {
		HTML string `json:"html"`
	} `json:"message"`
	DetailedMessage struct {
		HTML string `json:"html"`
	} `json:"detailedMessage"`
	Resource struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Web struct {
				Href string `json:"href"`
			} `json:"web"`
			SourceVersionDisplayURI struct {
				Href string `json:"href"`
			} `json:"sourceVersionDisplayUri"`
			Timeline struct {
				Href string `json:"href"`
			} `json:"timeline"`
			Badge struct {
				Href string `json:"href"`
			} `json:"badge"`
		} `json:"_links"`
		Properties struct {
		} `json:"properties"`
		Tags              []interface{} `json:"tags"`
		ValidationResults []interface{} `json:"validationResults"`
		Plans             []struct {
			PlanID string `json:"planId"`
		} `json:"plans"`
		TemplateParameters struct {
		} `json:"templateParameters"`
		TriggerInfo struct {
		} `json:"triggerInfo"`
		ID          int       `json:"id"`
		BuildNumber string    `json:"buildNumber"`
		Status      string    `json:"status"`
		Result      string    `json:"result"`
		QueueTime   time.Time `json:"queueTime"`
		StartTime   time.Time `json:"startTime"`
		FinishTime  time.Time `json:"finishTime"`
		URL         string    `json:"url"`
		Definition  struct {
			Links struct {
			} `json:"_links"`
			Drafts      []interface{} `json:"drafts"`
			ID          int           `json:"id"`
			Name        string        `json:"name"`
			URL         string        `json:"url"`
			URI         string        `json:"uri"`
			Path        string        `json:"path"`
			Type        string        `json:"type"`
			QueueStatus string        `json:"queueStatus"`
			Revision    int           `json:"revision"`
			Project     struct {
				ID             string    `json:"id"`
				Name           string    `json:"name"`
				URL            string    `json:"url"`
				State          string    `json:"state"`
				Revision       int       `json:"revision"`
				Visibility     string    `json:"visibility"`
				LastUpdateTime time.Time `json:"lastUpdateTime"`
			} `json:"project"`
		} `json:"definition"`
		Project struct {
			ID             string    `json:"id"`
			Name           string    `json:"name"`
			URL            string    `json:"url"`
			State          string    `json:"state"`
			Revision       int       `json:"revision"`
			Visibility     string    `json:"visibility"`
			LastUpdateTime time.Time `json:"lastUpdateTime"`
		} `json:"project"`
		URI           string `json:"uri"`
		SourceBranch  string `json:"sourceBranch"`
		SourceVersion string `json:"sourceVersion"`
		Queue         struct {
			Links struct {
			} `json:"_links"`
			ID   int    `json:"id"`
			Name string `json:"name"`
			Pool struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"pool"`
		} `json:"queue"`
		Priority     string `json:"priority"`
		Reason       string `json:"reason"`
		RequestedFor struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"requestedFor"`
		RequestedBy struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"requestedBy"`
		LastChangedDate time.Time `json:"lastChangedDate"`
		LastChangedBy   struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
			Descriptor string `json:"descriptor"`
		} `json:"lastChangedBy"`
		Parameters        string `json:"parameters"`
		OrchestrationPlan struct {
			PlanID string `json:"planId"`
		} `json:"orchestrationPlan"`
		Logs struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"logs"`
		Repository struct {
			ID                 string      `json:"id"`
			Type               string      `json:"type"`
			Name               string      `json:"name"`
			URL                string      `json:"url"`
			Clean              interface{} `json:"clean"`
			CheckoutSubmodules bool        `json:"checkoutSubmodules"`
		} `json:"repository"`
		KeepForever       bool        `json:"keepForever"`
		RetainedByRelease bool        `json:"retainedByRelease"`
		TriggeredByBuild  interface{} `json:"triggeredByBuild"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Server struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"server"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}

type Release struct {
	SubscriptionID string `json:"subscriptionId"`
	NotificationID int    `json:"notificationId"`
	ID             string `json:"id"`
	EventType      string `json:"eventType"`
	PublisherID    string `json:"publisherId"`
	Message        struct {
		HTML string `json:"html"`
	} `json:"message"`
	DetailedMessage struct {
		HTML string `json:"html"`
	} `json:"detailedMessage"`
	Resource struct {
		Environment struct {
			ID        int    `json:"id"`
			ReleaseID int    `json:"releaseId"`
			Name      string `json:"name"`
			Status    string `json:"status"`
			Variables struct {
			} `json:"variables"`
			VariableGroups     []interface{} `json:"variableGroups"`
			PreDeployApprovals []struct {
				ID               int       `json:"id"`
				Revision         int       `json:"revision"`
				ApprovalType     string    `json:"approvalType"`
				CreatedOn        time.Time `json:"createdOn"`
				ModifiedOn       time.Time `json:"modifiedOn"`
				Status           string    `json:"status"`
				Comments         string    `json:"comments"`
				IsAutomated      bool      `json:"isAutomated"`
				IsNotificationOn bool      `json:"isNotificationOn"`
				TrialNumber      int       `json:"trialNumber"`
				Attempt          int       `json:"attempt"`
				Rank             int       `json:"rank"`
				Release          struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"release"`
				ReleaseDefinition struct {
					ID               int         `json:"id"`
					Name             string      `json:"name"`
					Path             string      `json:"path"`
					ProjectReference interface{} `json:"projectReference"`
					URL              string      `json:"url"`
					Links            struct {
					} `json:"_links"`
				} `json:"releaseDefinition"`
				ReleaseEnvironment struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"releaseEnvironment"`
				URL string `json:"url"`
			} `json:"preDeployApprovals"`
			PostDeployApprovals []struct {
				ID               int       `json:"id"`
				Revision         int       `json:"revision"`
				ApprovalType     string    `json:"approvalType"`
				CreatedOn        time.Time `json:"createdOn"`
				ModifiedOn       time.Time `json:"modifiedOn"`
				Status           string    `json:"status"`
				Comments         string    `json:"comments"`
				IsAutomated      bool      `json:"isAutomated"`
				IsNotificationOn bool      `json:"isNotificationOn"`
				TrialNumber      int       `json:"trialNumber"`
				Attempt          int       `json:"attempt"`
				Rank             int       `json:"rank"`
				Release          struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"release"`
				ReleaseDefinition struct {
					ID               int         `json:"id"`
					Name             string      `json:"name"`
					Path             string      `json:"path"`
					ProjectReference interface{} `json:"projectReference"`
					URL              string      `json:"url"`
					Links            struct {
					} `json:"_links"`
				} `json:"releaseDefinition"`
				ReleaseEnvironment struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"releaseEnvironment"`
				URL string `json:"url"`
			} `json:"postDeployApprovals"`
			PreApprovalsSnapshot struct {
				Approvals []interface{} `json:"approvals"`
			} `json:"preApprovalsSnapshot"`
			PostApprovalsSnapshot struct {
				Approvals []interface{} `json:"approvals"`
			} `json:"postApprovalsSnapshot"`
			DeploySteps []struct {
				ID                  int    `json:"id"`
				DeploymentID        int    `json:"deploymentId"`
				Attempt             int    `json:"attempt"`
				Reason              string `json:"reason"`
				Status              string `json:"status"`
				OperationStatus     string `json:"operationStatus"`
				ReleaseDeployPhases []struct {
					ID             int    `json:"id"`
					PhaseID        string `json:"phaseId"`
					Name           string `json:"name"`
					Rank           int    `json:"rank"`
					PhaseType      string `json:"phaseType"`
					Status         string `json:"status"`
					RunPlanID      string `json:"runPlanId"`
					DeploymentJobs []struct {
						Job struct {
							ID               int           `json:"id"`
							TimelineRecordID string        `json:"timelineRecordId"`
							Name             string        `json:"name"`
							DateStarted      time.Time     `json:"dateStarted"`
							DateEnded        time.Time     `json:"dateEnded"`
							StartTime        time.Time     `json:"startTime"`
							FinishTime       time.Time     `json:"finishTime"`
							Status           string        `json:"status"`
							Rank             int           `json:"rank"`
							Issues           []interface{} `json:"issues"`
							AgentName        string        `json:"agentName"`
						} `json:"job"`
						Tasks []struct {
							ID               int           `json:"id"`
							TimelineRecordID string        `json:"timelineRecordId"`
							Name             string        `json:"name"`
							DateStarted      time.Time     `json:"dateStarted"`
							DateEnded        time.Time     `json:"dateEnded"`
							StartTime        time.Time     `json:"startTime"`
							FinishTime       time.Time     `json:"finishTime"`
							Status           string        `json:"status"`
							Rank             int           `json:"rank"`
							Issues           []interface{} `json:"issues"`
							AgentName        string        `json:"agentName"`
							Task             struct {
								ID      string `json:"id"`
								Name    string `json:"name"`
								Version string `json:"version"`
							} `json:"task,omitempty"`
						} `json:"tasks"`
					} `json:"deploymentJobs"`
					ManualInterventions []interface{} `json:"manualInterventions"`
					StartedOn           time.Time     `json:"startedOn"`
				} `json:"releaseDeployPhases"`
				RequestedBy struct {
					DisplayName string `json:"displayName"`
					URL         string `json:"url"`
					Links       struct {
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"_links"`
					ID         string `json:"id"`
					UniqueName string `json:"uniqueName"`
					ImageURL   string `json:"imageUrl"`
					Descriptor string `json:"descriptor"`
				} `json:"requestedBy"`
				RequestedFor struct {
					DisplayName string `json:"displayName"`
					URL         string `json:"url"`
					Links       struct {
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"_links"`
					ID         string `json:"id"`
					UniqueName string `json:"uniqueName"`
					ImageURL   string `json:"imageUrl"`
					Descriptor string `json:"descriptor"`
				} `json:"requestedFor"`
				QueuedOn       time.Time `json:"queuedOn"`
				LastModifiedBy struct {
					DisplayName string `json:"displayName"`
					URL         string `json:"url"`
					Links       struct {
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"_links"`
					ID         string `json:"id"`
					UniqueName string `json:"uniqueName"`
					ImageURL   string `json:"imageUrl"`
					Descriptor string `json:"descriptor"`
				} `json:"lastModifiedBy"`
				LastModifiedOn time.Time     `json:"lastModifiedOn"`
				HasStarted     bool          `json:"hasStarted"`
				Tasks          []interface{} `json:"tasks"`
				RunPlanID      string        `json:"runPlanId"`
				Issues         []interface{} `json:"issues"`
			} `json:"deploySteps"`
			Rank                    int `json:"rank"`
			DefinitionEnvironmentID int `json:"definitionEnvironmentId"`
			EnvironmentOptions      struct {
				EmailNotificationType        string `json:"emailNotificationType"`
				EmailRecipients              string `json:"emailRecipients"`
				SkipArtifactsDownload        bool   `json:"skipArtifactsDownload"`
				TimeoutInMinutes             int    `json:"timeoutInMinutes"`
				EnableAccessToken            bool   `json:"enableAccessToken"`
				PublishDeploymentStatus      bool   `json:"publishDeploymentStatus"`
				BadgeEnabled                 bool   `json:"badgeEnabled"`
				AutoLinkWorkItems            bool   `json:"autoLinkWorkItems"`
				PullRequestDeploymentEnabled bool   `json:"pullRequestDeploymentEnabled"`
			} `json:"environmentOptions"`
			Demands    []interface{} `json:"demands"`
			Conditions []struct {
				Result        bool   `json:"result"`
				Name          string `json:"name"`
				ConditionType string `json:"conditionType"`
				Value         string `json:"value"`
			} `json:"conditions"`
			CreatedOn            time.Time     `json:"createdOn"`
			ModifiedOn           time.Time     `json:"modifiedOn"`
			WorkflowTasks        []interface{} `json:"workflowTasks"`
			DeployPhasesSnapshot []struct {
				DeploymentInput struct {
					ParallelExecution struct {
						ParallelExecutionType string `json:"parallelExecutionType"`
					} `json:"parallelExecution"`
					AgentSpecification     interface{} `json:"agentSpecification"`
					SkipArtifactsDownload  bool        `json:"skipArtifactsDownload"`
					ArtifactsDownloadInput struct {
						DownloadInputs []interface{} `json:"downloadInputs"`
					} `json:"artifactsDownloadInput"`
					QueueID                   int           `json:"queueId"`
					Demands                   []interface{} `json:"demands"`
					EnableAccessToken         bool          `json:"enableAccessToken"`
					TimeoutInMinutes          int           `json:"timeoutInMinutes"`
					JobCancelTimeoutInMinutes int           `json:"jobCancelTimeoutInMinutes"`
					Condition                 string        `json:"condition"`
					OverrideInputs            struct {
					} `json:"overrideInputs"`
				} `json:"deploymentInput"`
				Rank          int           `json:"rank"`
				PhaseType     string        `json:"phaseType"`
				Name          string        `json:"name"`
				RefName       interface{}   `json:"refName"`
				WorkflowTasks []interface{} `json:"workflowTasks"`
			} `json:"deployPhasesSnapshot"`
			Owner struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
				Descriptor string `json:"descriptor"`
			} `json:"owner"`
			Schedules []interface{} `json:"schedules"`
			Release   struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				URL   string `json:"url"`
				Links struct {
					Web struct {
						Href string `json:"href"`
					} `json:"web"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
			} `json:"release"`
			ReleaseDefinition struct {
				ID               int         `json:"id"`
				Name             string      `json:"name"`
				Path             string      `json:"path"`
				ProjectReference interface{} `json:"projectReference"`
				URL              string      `json:"url"`
				Links            struct {
					Web struct {
						Href string `json:"href"`
					} `json:"web"`
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
			} `json:"releaseDefinition"`
			ReleaseCreatedBy struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
				Descriptor string `json:"descriptor"`
			} `json:"releaseCreatedBy"`
			TriggerReason     string  `json:"triggerReason"`
			TimeToDeploy      float64 `json:"timeToDeploy"`
			ProcessParameters struct {
			} `json:"processParameters"`
			PreDeploymentGatesSnapshot struct {
				ID           int           `json:"id"`
				GatesOptions interface{}   `json:"gatesOptions"`
				Gates        []interface{} `json:"gates"`
			} `json:"preDeploymentGatesSnapshot"`
			PostDeploymentGatesSnapshot struct {
				ID           int           `json:"id"`
				GatesOptions interface{}   `json:"gatesOptions"`
				Gates        []interface{} `json:"gates"`
			} `json:"postDeploymentGatesSnapshot"`
		} `json:"environment"`
		Project struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"project"`
		Deployment struct {
			ID      int `json:"id"`
			Release struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				URL       string `json:"url"`
				Artifacts []struct {
					SourceID            string `json:"sourceId"`
					Type                string `json:"type"`
					Alias               string `json:"alias"`
					DefinitionReference struct {
						ArtifactSourceDefinitionURL struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"artifactSourceDefinitionUrl"`
						Branches struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"branches"`
						BuildURI struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"buildUri"`
						Definition struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"definition"`
						IsMultiDefinitionType struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"IsMultiDefinitionType"`
						IsXamlBuildArtifactType struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"IsXamlBuildArtifactType"`
						Project struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"project"`
						RepositoryProvider struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"repository.provider"`
						Repository struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"repository"`
						RequestedFor struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"requestedFor"`
						RequestedForID struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"requestedForId"`
						SourceVersion struct {
							ID   string      `json:"id"`
							Name interface{} `json:"name"`
						} `json:"sourceVersion"`
						Version struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"version"`
						ArtifactSourceVersionURL struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"artifactSourceVersionUrl"`
						DefaultVersionType struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"defaultVersionType"`
						Branch struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"branch"`
					} `json:"definitionReference"`
					IsPrimary  bool `json:"isPrimary,omitempty"`
					IsRetained bool `json:"isRetained"`
				} `json:"artifacts"`
				WebAccessURI string `json:"webAccessUri"`
				Links        struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Web struct {
						Href string `json:"href"`
					} `json:"web"`
					Logs struct {
						Href string `json:"href"`
					} `json:"logs"`
				} `json:"_links"`
				Description string `json:"description"`
				Reason      string `json:"reason"`
			} `json:"release"`
			ReleaseDefinition struct {
				ID               int    `json:"id"`
				Name             string `json:"name"`
				Path             string `json:"path"`
				ProjectReference struct {
					ID   string      `json:"id"`
					Name interface{} `json:"name"`
				} `json:"projectReference"`
				URL   string `json:"url"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Web struct {
						Href string `json:"href"`
					} `json:"web"`
				} `json:"_links"`
			} `json:"releaseDefinition"`
			ReleaseEnvironment struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				URL   string `json:"url"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Web struct {
						Href string `json:"href"`
					} `json:"web"`
				} `json:"_links"`
			} `json:"releaseEnvironment"`
			ProjectReference        interface{} `json:"projectReference"`
			DefinitionEnvironmentID int         `json:"definitionEnvironmentId"`
			Attempt                 int         `json:"attempt"`
			Reason                  string      `json:"reason"`
			DeploymentStatus        string      `json:"deploymentStatus"`
			OperationStatus         string      `json:"operationStatus"`
			RequestedBy             struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
				Descriptor string `json:"descriptor"`
			} `json:"requestedBy"`
			RequestedFor struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
				Descriptor string `json:"descriptor"`
			} `json:"requestedFor"`
			QueuedOn       time.Time `json:"queuedOn"`
			StartedOn      time.Time `json:"startedOn"`
			CompletedOn    time.Time `json:"completedOn"`
			LastModifiedOn time.Time `json:"lastModifiedOn"`
			LastModifiedBy struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
				Descriptor string `json:"descriptor"`
			} `json:"lastModifiedBy"`
			Conditions []struct {
				Name          string `json:"name"`
				ConditionType string `json:"conditionType"`
				Value         string `json:"value"`
			} `json:"conditions"`
			PreDeployApprovals []struct {
				ID               int       `json:"id"`
				Revision         int       `json:"revision"`
				ApprovalType     string    `json:"approvalType"`
				CreatedOn        time.Time `json:"createdOn"`
				ModifiedOn       time.Time `json:"modifiedOn"`
				Status           string    `json:"status"`
				Comments         string    `json:"comments"`
				IsAutomated      bool      `json:"isAutomated"`
				IsNotificationOn bool      `json:"isNotificationOn"`
				TrialNumber      int       `json:"trialNumber"`
				Attempt          int       `json:"attempt"`
				Rank             int       `json:"rank"`
				Release          struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"release"`
				ReleaseDefinition struct {
					ID               int         `json:"id"`
					Name             string      `json:"name"`
					ProjectReference interface{} `json:"projectReference"`
					URL              string      `json:"url"`
					Links            struct {
					} `json:"_links"`
				} `json:"releaseDefinition"`
				ReleaseEnvironment struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"releaseEnvironment"`
				URL string `json:"url"`
			} `json:"preDeployApprovals"`
			PostDeployApprovals []struct {
				ID               int       `json:"id"`
				Revision         int       `json:"revision"`
				ApprovalType     string    `json:"approvalType"`
				CreatedOn        time.Time `json:"createdOn"`
				ModifiedOn       time.Time `json:"modifiedOn"`
				Status           string    `json:"status"`
				Comments         string    `json:"comments"`
				IsAutomated      bool      `json:"isAutomated"`
				IsNotificationOn bool      `json:"isNotificationOn"`
				TrialNumber      int       `json:"trialNumber"`
				Attempt          int       `json:"attempt"`
				Rank             int       `json:"rank"`
				Release          struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"release"`
				ReleaseDefinition struct {
					ID               int         `json:"id"`
					Name             string      `json:"name"`
					ProjectReference interface{} `json:"projectReference"`
					URL              string      `json:"url"`
					Links            struct {
					} `json:"_links"`
				} `json:"releaseDefinition"`
				ReleaseEnvironment struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					URL   string `json:"url"`
					Links struct {
					} `json:"_links"`
				} `json:"releaseEnvironment"`
				URL string `json:"url"`
			} `json:"postDeployApprovals"`
			Links struct {
			} `json:"_links"`
		} `json:"deployment"`
		Comment interface{} `json:"comment"`
		Data    struct {
			ReleaseProperties          string        `json:"releaseProperties"`
			EnvironmentStatuses        string        `json:"environmentStatuses"`
			WorkItems                  []interface{} `json:"workItems"`
			PreviousReleaseEnvironment struct {
				Status string `json:"status"`
			} `json:"previousReleaseEnvironment"`
			Commits     []interface{} `json:"commits"`
			TestResults struct {
			} `json:"testResults"`
			MoreWorkItemsMessage string `json:"moreWorkItemsMessage"`
		} `json:"data"`
		StageName string `json:"stageName"`
		AttemptID int    `json:"attemptId"`
		ID        int    `json:"id"`
		URL       string `json:"url"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Server struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"server"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}
