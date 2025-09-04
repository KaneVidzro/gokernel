package mail

type EmailJob struct {
    To      string
    Subject string
    Body    string
}

func (m *MailService) Queue(job EmailJob) {
    go func() {
        _ = m.Send(job.To, job.Subject, job.Body)
    }()
}
