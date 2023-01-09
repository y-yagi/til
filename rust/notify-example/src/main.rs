use notify::{Config, RecommendedWatcher, RecursiveMode, Watcher};
use std::path::Path;
use std::time::Duration;

fn main() {
    let path = std::env::args()
        .nth(1)
        .expect("Argument 1 needs to be a path");
    println!("watching {}", path);
    if let Err(e) = watch(path) {
        println!("error: {:?}", e)
    }
}

fn watch<P: AsRef<Path>>(path: P) -> notify::Result<()> {
    let (watcher_tx, watcher_rx) = crossbeam_channel::unbounded();
    let (sig_tx, sig_rx) = crossbeam_channel::unbounded();

    let mut watcher = RecommendedWatcher::new(watcher_tx, Config::default())?;
    watcher.watch(path.as_ref(), RecursiveMode::Recursive)?;

    ctrlc::set_handler(move || sig_tx.send(()).expect("could not send signal on channel."))
        .expect("error setting Ctrl-C handler");

    loop {
        crossbeam_channel::select! {
            recv(watcher_rx) -> res => {
              match res {
                  Ok(event) => println!("changed: {:?}", event),
                  Err(e) => println!("watch error: {:?}", e),
              }
            }
            recv(sig_rx) -> _res => {
                println!("got signal");
                break
            }
            default(Duration::from_secs(10)) => {
                println!("timeout");
                break
            }
        }
    }

    println!("watch end");
    Ok(())
}
